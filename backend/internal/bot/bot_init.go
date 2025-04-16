package bot

import (
	"log"

	"github.com/bullockz21/beer_bot/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(cfg *configs.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}
	log.Println("Bot успешно инициализирован")
	return bot, nil
}
