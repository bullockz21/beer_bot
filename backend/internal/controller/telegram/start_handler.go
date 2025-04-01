// internal/telegram/start_handler.go
package telegram

import (
	"context"
	"log"

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
// Данные пользователя извлекаются напрямую из update.Message.
func (h *StartHandler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	telegramID := update.Message.From.ID
	username := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	language := update.Message.From.LanguageCode

	// Вызываем usecase для создания пользователя.
	if _, err := h.userUC.CreateUser(ctx, telegramID, username, firstName, language); err != nil {
		log.Printf("[ERROR] HandleStart: %v", err)
		h.userPresenter.PresentError(update.Message.Chat.ID, "Не удалось создать пользователя")
		return
	}

	// Отправляем приветственное сообщение.
	if err := h.userPresenter.PresentWelcomeMessage(update.Message.Chat.ID, firstName); err != nil {
		log.Printf("Ошибка отправки приветственного сообщения: %v", err)
	}
}
