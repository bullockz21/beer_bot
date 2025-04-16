package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handler объединяет обработчики команд и callback.
type Handler struct {
	bot             *tgbotapi.BotAPI
	commandHandler  *CommandHandler
	callbackHandler *CallbackHandler
}

// NewHandler создает новый экземпляр Handler.
func NewHandler(bot *tgbotapi.BotAPI, commandHandler *CommandHandler, callbackHandler *CallbackHandler) *Handler {
	return &Handler{
		bot:             bot,
		commandHandler:  commandHandler,
		callbackHandler: callbackHandler,
	}
}

// ProcessUpdate распределяет обновления по типам.
func (h *Handler) ProcessUpdate(ctx context.Context, update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		h.commandHandler.HandleCommand(ctx, update)
	} else if update.CallbackQuery != nil {
		h.callbackHandler.HandleCallback(ctx, update)
	}
}
