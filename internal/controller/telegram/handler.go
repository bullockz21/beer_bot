package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/bullockz21/beer_bot/internal/entity"
	"github.com/bullockz21/beer_bot/internal/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot    *tgbotapi.BotAPI
	userUC *usecase.UserUseCase
}

func NewHandler(bot *tgbotapi.BotAPI, userUC *usecase.UserUseCase) *Handler {
	return &Handler{
		bot:    bot,
		userUC: userUC,
	}
}

func (h *Handler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	tgUser := entity.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		Language:   update.Message.From.LanguageCode,
	}

	user, err := h.userUC.HandleStart(ctx, &tgUser)
	if err != nil {
		log.Printf("Error handling /start: %v", err)
		msg.Text = "🚫 Произошла ошибка, попробуйте позже"
	} else {
		msg.Text = fmt.Sprintf(
			"👋 Привет, %s!\n\n"+
				"Ваш аккаунт успешно зарегистрирован!\n"+
				"ID: %d\n"+
				"Username: @%s",
			user.FirstName,
			user.ID,
			user.Username,
		)
	}

	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
