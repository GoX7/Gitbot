package handlers

import telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewCallback(bot *telegram.BotAPI, update telegram.Update) {
	message := telegram.NewMessage(update.CallbackQuery.Message.Chat.ID, "callback not valide 🤨")
	bot.Send(message)
}
