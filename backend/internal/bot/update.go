package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ListenUpdates получает обновления от Telegram и передает их обработчику
func ListenUpdates(ctx context.Context, bot *tgbotapi.BotAPI, handler *Handler) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handler.ProcessUpdate(ctx, update)
	}
}
