package telegram

import (
	"context"
	"log"

	dtoUser "github.com/bullockz21/beer_bot/internal/dto/user"
	presenterUser "github.com/bullockz21/beer_bot/internal/presenter/user"
	usecaseUser "github.com/bullockz21/beer_bot/internal/usecase/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// StartHandler обрабатывает команду /start.
type StartHandler struct {
	userUC        *usecaseUser.UserUseCase
	userPresenter *presenterUser.UserPresenter
}

// NewStartHandler создает новый StartHandler.
func NewStartHandler(userUC *usecaseUser.UserUseCase, userPresenter *presenterUser.UserPresenter) *StartHandler {
	return &StartHandler{
		userUC:        userUC,
		userPresenter: userPresenter,
	}
}

// HandleStart обрабатывает команду /start.
func (h *StartHandler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	req := dtoUser.UserCreateRequest{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		Language:   update.Message.From.LanguageCode,
	}

	if _, err := h.userUC.HandleStart(ctx, &req); err != nil {
		log.Printf("[ERROR] HandleStart: %v", err)
		h.userPresenter.PresentError(update.Message.Chat.ID, "Не удалось создать пользователя")
		return
	}

	// Отправляем приветственное сообщение
	if err := h.userPresenter.PresentWelcomeMessage(update.Message.Chat.ID, update.Message.From.FirstName); err != nil {
		log.Printf("Ошибка отправки приветственного сообщения: %v", err)
	}
}
