package bot

import (
	"context"

	"github.com/bullockz21/beer_bot/internal/controller/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ListenUpdates получает обновления от Telegram и передает их обработчику.
func ListenUpdates(ctx context.Context, bot *tgbotapi.BotAPI, handler *telegram.Handler) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handler.ProcessUpdate(ctx, update)
	}
}
