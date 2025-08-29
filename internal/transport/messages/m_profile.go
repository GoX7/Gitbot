package messages

import (
	"fmt"
	"regexp"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gox7/gitbot/internal/github"
	"github.com/gox7/gitbot/internal/services"
	"github.com/gox7/gitbot/models"
)

func SendProfile(bot *telegram.BotAPI, users *services.PostgresUsers, redisDB *services.Redis, message *telegram.Message) {
	status, err := redisDB.Get(fmt.Sprint(message.From.ID))
	if err != nil {
		redisDB.Set(fmt.Sprint(message.From.ID), "")
	}

	if status == "profile:register" {
		if message.Text == "❌ Cancel" {
			text := "❌ Registration canceled.\n\n" +
				"You can always start again by selecting *📈 My GitHub*."
			response := telegram.NewMessage(message.Chat.ID, text)
			response.ParseMode = "Markdown"
			response.ReplyMarkup = buttons()

			redisDB.Set(fmt.Sprint(message.From.ID), "")
			bot.Send(response)
			return
		}

		if !checkName(bot, message) {
			return
		}
		if !checkAccout(bot, message) {
			redisDB.Set(fmt.Sprint(message.From.ID), "")
			return
		}
		if !registerUser(bot, users, message) {
			redisDB.Set(fmt.Sprint(message.From.ID), "")
			return
		}

		text := "✅ Registration complete!\n\n" +
			"Your GitHub account has been linked successfully.\n\n" +
			"Use the menu below to explore your stats."
		response := telegram.NewMessage(message.Chat.ID, text)
		response.ReplyMarkup = buttons()

		redisDB.Set(fmt.Sprint(message.From.ID), "")
		bot.Send(response)
	} else {
		model, err := users.Search(message.From.ID)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				redisDB.Set(fmt.Sprint(message.From.ID), "profile:register")
				responseRegister(bot, message)
				return
			}

			response := telegram.NewMessage(message.Chat.ID, "⚠️ Server error.\n\nPlease try again later.")
			response.ReplyMarkup = buttons()

			bot.Send(response)
		} else {
			responseLogin(bot, message, model)
		}
	}
}

func checkName(bot *telegram.BotAPI, message *telegram.Message) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9-]{1,39}$")
	if !re.MatchString(message.Text) {
		text := "🚫 Invalid username.\n\n" +
			"GitHub usernames can only contain *letters, numbers and dashes* (1–39 characters)."
		response := telegram.NewMessage(message.Chat.ID, text)
		response.ParseMode = "Markdown"

		bot.Send(response)
		return false
	}
	return true
}

func checkAccout(bot *telegram.BotAPI, message *telegram.Message) bool {
	status, err := github.CheckAccout(message.Text)
	if err != nil {
		text := "⚠️ GitHub API error.\n\n" +
			"Please try again later."
		response := telegram.NewMessage(message.Chat.ID, text)
		response.ReplyMarkup = buttons()

		bot.Send(response)
		return false
	}

	if !status {
		response := telegram.NewMessage(message.Chat.ID, "🚫 GitHub account not found.\n\nPlease check the username and try again.")
		response.ReplyMarkup = buttons()

		bot.Send(response)
		return false
	}

	return true
}

func registerUser(bot *telegram.BotAPI, users *services.PostgresUsers, message *telegram.Message) bool {
	if err := users.Register(message.From.ID, message.From.UserName, message.Text); err != nil {
		response := telegram.NewMessage(message.Chat.ID, "⚠️ Server error.\n\nPlease try again later.")
		bot.Send(response)
		return false
	}
	return true
}

func buttons() telegram.ReplyKeyboardMarkup {
	return telegram.NewReplyKeyboard(
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("🔍 Find profile"),
			telegram.NewKeyboardButton("📈 My GitHub"),
		),
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("ℹ️ about"),
		),
	)
}

func responseLogin(bot *telegram.BotAPI, message *telegram.Message, model *models.UserOrm) {
	account, err := github.GetAccout(model.Github)
	if err != nil {
		if err.Error() == "ANF" {
			response := telegram.NewMessage(message.Chat.ID, "🚫 GitHub account not found.\n\nPlease later and try again.")
			bot.Send(response)
			return
		}

		text := "⚠️ GitHub API error.\n\n" +
			"Please try again later."
		response := telegram.NewMessage(message.Chat.ID, text)
		bot.Send(response)
		return
	}

	stars, forks := github.GetStarsForks(model.Github)
	caption := fmt.Sprintf("🖼 *%s* (%s)\n\n", account.Login, account.Name)
	if account.Bio != "" {
		caption = caption + fmt.Sprintf("💬 _%s_\n\n", account.Bio)
	}
	if account.Email != "" {
		caption = caption + fmt.Sprintf("✉️ Email: %s\n", account.Email)
	}
	if account.Location != "" {
		caption = caption + fmt.Sprintf("📍 Location: %s\n", account.Location)
	}

	caption = caption + fmt.Sprintf("📦 Repositories: %d  ⭐ Stars: %d  🍴 Forks: %d\n"+
		"👥 Followers: %d  🔗 Following: %d\n"+
		"🌐 Profile: [GitHub](https://github.com/%s)\n\n"+
		"🔥 Keep coding and pushing commits! 🚀",
		account.PublicRepos, stars, forks,
		account.Followers, account.Following,
		account.Login,
	)

	response := telegram.NewPhoto(message.Chat.ID, telegram.FileURL(account.AvatarURL))
	response.Caption = caption
	response.ParseMode = "markdown"

	bot.Send(response)
}

func responseRegister(bot *telegram.BotAPI, message *telegram.Message) {
	text := "👋 It looks like you are not registered yet.\n\n" +
		"Please send me your GitHub username to continue.\n\n" +
		"Example: *torvalds*"
	response := telegram.NewMessage(message.Chat.ID, text)
	response.ParseMode = "Markdown"
	buttons := telegram.NewReplyKeyboard(
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("❌ Cancel"),
		),
	)

	response.ReplyMarkup = buttons
	bot.Send(response)
}
