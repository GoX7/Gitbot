package messages

import (
	"fmt"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gox7/gitbot/internal/github"
	"github.com/gox7/gitbot/internal/services"
)

func SendSearch(bot *telegram.BotAPI, redisDB *services.Redis, message *telegram.Message) {
	status, err := redisDB.Get(fmt.Sprint(message.From.ID))
	if err != nil {
		redisDB.Set(fmt.Sprint(message.From.ID), "")
	}

	if status == "search:username" {
		if message.Text == "❌ Cancel" {
			text := "❌ Search canceled.\n\n" +
				"You can always start again by selecting *🔍 Find profile*."
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

		account, err := github.GetAccout(message.Text)
		if err != nil {
			text := "⚠️ GitHub API error.\n\n" +
				"Please try again later."
			response := telegram.NewMessage(message.Chat.ID, text)
			bot.Send(response)
			redisDB.Set(fmt.Sprint(message.From.ID), "")
			return
		}

		stars, forks := github.GetStarsForks(message.Text)
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

		response.ReplyMarkup = buttons()
		redisDB.Set(fmt.Sprint(message.From.ID), "")
		bot.Send(response)

	} else {
		redisDB.Set(fmt.Sprint(message.From.ID), "search:username")
		responseSearchIntro(bot, message)
	}
}

func responseSearchIntro(bot *telegram.BotAPI, message *telegram.Message) {
	text := "🔍 Please enter the GitHub username you want to search.\n\n" +
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
