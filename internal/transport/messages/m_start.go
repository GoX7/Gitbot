package messages

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendStart(bot *telegram.BotAPI, message *telegram.Message) {
	text := "📊 GitHub Analyzer Bot\n\n" +
		"This bot helps you quickly get statistics about any GitHub profile\n"

	response := telegram.NewMessage(message.Chat.ID, text)
	buttons := telegram.NewReplyKeyboard(
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("🔍 Find profile"),
			telegram.NewKeyboardButton("📈 My GitHub"),
		),
		telegram.NewKeyboardButtonRow(
			telegram.NewKeyboardButton("ℹ️ about"),
		),
	)

	response.ReplyMarkup = buttons
	bot.Send(response)
}
