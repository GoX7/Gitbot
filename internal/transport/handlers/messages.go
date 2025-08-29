package handlers

import (
	"fmt"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gox7/gitbot/internal/services"
	"github.com/gox7/gitbot/internal/transport/messages"
)

var (
	postgresUsers *services.PostgresUsers
	redis         *services.Redis
)

func Resource(users *services.PostgresUsers, redisDB *services.Redis) {
	postgresUsers = users
	redis = redisDB
}

func NewMessage(bot *telegram.BotAPI, update telegram.Update) {
	status, err := redis.Get(fmt.Sprint(update.Message.From.ID))
	if err == nil {
		splits := strings.Split(status, ":")
		if splits[0] == "profile" {
			messages.SendProfile(bot, postgresUsers, redis, update.Message)
			return
		} else if splits[0] == "search" {
			messages.SendSearch(bot, redis, update.Message)
			return
		}
	}

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			messages.SendStart(bot, update.Message)
		case "search":
			messages.SendSearch(bot, redis, update.Message)
		case "profile":
			messages.SendProfile(bot, postgresUsers, redis, update.Message)
		case "about":
			messages.SendAbout(bot, update.Message)
		}
	}

	switch update.Message.Text {
	case "🔍 Find profile":
		messages.SendSearch(bot, redis, update.Message)
	case "📈 My GitHub":
		messages.SendProfile(bot, postgresUsers, redis, update.Message)
	case "ℹ️ about":
		messages.SendAbout(bot, update.Message)
	}
}
