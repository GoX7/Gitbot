package messages

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendAbout(bot *telegram.BotAPI, message *telegram.Message) {
	text := "ℹ️ GitHub Analyzer Bot\n\n" +
		"This bot allows you to analyze any GitHub profile\n\n" +
		"📌 Author: gox7\n" +
		"💻 Written in Go"
	response := telegram.NewMessage(message.Chat.ID, text)
	buttons := telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonURL("🌐 github", "https://github.com/gox7"),
		),
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonURL("📄 telegram", "https://t.me/sex1nk"),
		),
	)

	response.ReplyMarkup = buttons
	bot.Send(response)
}
