package middleware

import (
	"fmt"
	"log/slog"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewLogger(logger *slog.Logger, update telegram.Update) {
	if update.Message != nil {
		logger.Info(
			"message",
			slog.String("userid", fmt.Sprint(update.Message.From.ID)),
			slog.String("username", update.Message.From.UserName),
			slog.String("data", update.Message.Text),
		)
	} else if update.CallbackQuery != nil {
		logger.Info(
			"callback",
			slog.String("userid", fmt.Sprint(update.Message.From.ID)),
			slog.String("username", update.Message.From.UserName),
			slog.String("data", update.CallbackQuery.Data),
		)
	} else {
		logger.Info(
			"unkcown",
			slog.String("userid", fmt.Sprint(update.Message.From.ID)),
			slog.String("username", update.Message.From.UserName),
		)
	}
}
