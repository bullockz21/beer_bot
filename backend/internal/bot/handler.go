package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot             *tgbotapi.BotAPI
	commandHandler  *CommandHandler
	callbackHandler *CallbackHandler
}

func NewHandler(bot *tgbotapi.BotAPI, commandHandler *CommandHandler, callbackHandler *CallbackHandler) *Handler {
	return &Handler{
		bot:             bot,
		commandHandler:  commandHandler,
		callbackHandler: callbackHandler,
	}
}

func (h *Handler) ProcessUpdate(ctx context.Context, update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		h.commandHandler.HandleCommand(ctx, update)
	} else if update.CallbackQuery != nil {
		h.callbackHandler.HandleCallback(ctx, update)
	}
}
