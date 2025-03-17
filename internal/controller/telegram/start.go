package telegram

import (
	"context"
	"log"

	dto "github.com/bullockz21/beer_bot/internal/dto/user"
	presenter "github.com/bullockz21/beer_bot/internal/presenter/user"
	resource "github.com/bullockz21/beer_bot/internal/resource/user"
	usecase "github.com/bullockz21/beer_bot/internal/usecase/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot          *tgbotapi.BotAPI
	userUC       *usecase.UserUseCase
	presenter    *presenter.UserPresenter // Добавлено явное поле
	userResource *resource.UserResource   // Добавлено поле для ресурса
}

func NewHandler(
	bot *tgbotapi.BotAPI,
	userUC *usecase.UserUseCase,
	presenter *presenter.UserPresenter,
	userResource *resource.UserResource,
) *Handler {
	return &Handler{
		bot:          bot,
		userUC:       userUC,
		presenter:    presenter,
		userResource: userResource,
	}
}

func (h *Handler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	// Создаем DTO из входящих данных
	req := dto.UserCreateRequest{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		Language:   update.Message.From.LanguageCode,
	}

	// Вызываем Use Case
	user, err := h.userUC.HandleStart(ctx, &req)
	if err != nil {
		log.Printf("[ERROR] HandleStart: %v", err)
		_ = h.presenter.PresentError(update.Message.Chat.ID, "Не удалось создать пользователя")
		return
	}

	// Преобразуем сущность в DTO ответа
	response := h.userResource.ToResponse(user)

	// Используем презентер для успешного ответа
	if err := h.presenter.PresentUserSuccess(update.Message.Chat.ID, response); err != nil {
		log.Printf("Failed to send success message: %v", err)
	}
}
