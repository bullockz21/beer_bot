package telegram

import (
	"context"

	presenterUser "github.com/bullockz21/beer_bot/internal/presenter/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CommandHandler отвечает за обработку текстовых команд.
type CommandHandler struct {
	startHandler  *StartHandler
	userPresenter *presenterUser.UserPresenter
}

// NewCommandHandler создает новый CommandHandler.
func NewCommandHandler(startHandler *StartHandler, userPresenter *presenterUser.UserPresenter) *CommandHandler {
	return &CommandHandler{
		startHandler:  startHandler,
		userPresenter: userPresenter,
	}
}

// HandleCommand распределяет команды.
func (h *CommandHandler) HandleCommand(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		h.startHandler.HandleStart(ctx, update)
	default:
		h.userPresenter.PresentError(update.Message.Chat.ID, "Неизвестная команда.")
	}
}
