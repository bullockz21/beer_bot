package telegram

import (
	"context"
	"log"

	dtoUser "github.com/bullockz21/beer_bot/internal/dto/user"
	presenterUser "github.com/bullockz21/beer_bot/internal/presenter/user"
	resourceUser "github.com/bullockz21/beer_bot/internal/resource/user"
	usecaseUser "github.com/bullockz21/beer_bot/internal/usecase/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CommandHandler отвечает за обработку текстовых команд.
type CommandHandler struct {
	userUC        *usecaseUser.UserUseCase
	userPresenter *presenterUser.UserPresenter
	userResource  *resourceUser.UserResource
}

// NewCommandHandler создает новый CommandHandler.
func NewCommandHandler(userUC *usecaseUser.UserUseCase, userPresenter *presenterUser.UserPresenter, userResource *resourceUser.UserResource) *CommandHandler {
	return &CommandHandler{
		userUC:        userUC,
		userPresenter: userPresenter,
		userResource:  userResource,
	}
}

// HandleCommand распределяет команды.
func (h *CommandHandler) HandleCommand(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		h.HandleStart(ctx, update)
	default:
		h.userPresenter.PresentError(update.Message.Chat.ID, "Неизвестная команда.")
	}
}

// HandleStart обрабатывает команду /start.
func (h *CommandHandler) HandleStart(ctx context.Context, update tgbotapi.Update) {
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

	// Отправляем приветственное сообщение с инлайн-клавиатурой
	if err := h.userPresenter.PresentWelcomeMessage(update.Message.Chat.ID, update.Message.From.FirstName); err != nil {
		log.Printf("Ошибка отправки приветственного сообщения: %v", err)
	}
}
