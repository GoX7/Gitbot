package transport

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gox7/gitbot/internal/services"
	"github.com/gox7/gitbot/internal/transport/handlers"
	"github.com/gox7/gitbot/internal/transport/middleware"
	"github.com/gox7/gitbot/models"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	redis         *services.Redis
	postgresUsers *services.PostgresUsers
	serverLogger  *slog.Logger
)

func Resource(logger *slog.Logger, users *services.PostgresUsers, redisDB *services.Redis) {
	redis = redisDB
	postgresUsers = users
	serverLogger = logger
}

func Register(config *models.Config) {
	bot, err := telegram.NewBotAPI(config.TELEGRAM_TOKEN)
	if err != nil {
		fmt.Println("[-] telegram.bot: " + err.Error())
		os.Exit(1)
	}

	updateCFG := telegram.NewUpdate(0)
	updateCFG.Timeout = 30

	handlers.Resource(postgresUsers, redis)

	fmt.Println("[+] telegram.bot: updates...")
	updates := bot.GetUpdatesChan(updateCFG)
	for update := range updates {
		go middleware.NewLogger(serverLogger, update)
		if update.Message != nil {
			go handlers.NewMessage(bot, update)
		} else if update.CallbackQuery != nil {
			go handlers.NewCallback(bot, update)
		}
	}
}
