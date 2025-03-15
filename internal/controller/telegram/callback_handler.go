package telegram

import (
	"context"
	"log"

	"github.com/bullockz21/beer_bot/internal/presenter/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler struct {
	bot *tgbotapi.BotAPI
}

func NewCallbackHandler(bot *tgbotapi.BotAPI) *CallbackHandler {
	return &CallbackHandler{bot: bot}
}

func (h *CallbackHandler) HandleCallback(ctx context.Context, update tgbotapi.Update) {
	callback := update.CallbackQuery
	chatID := callback.Message.Chat.ID
	data := callback.Data

	var text string
	switch data {
	case buttons.MenuButton.Data:
		newKeyboard := buttons.InlineKeyboardColumn(buttons.ShawarmaButton, buttons.DrinksButton, buttons.DessertsButton, buttons.BackButton)
		// Создаем объект редактирования клавиатуры. chatID и messageID берутся из callback.
		edit := tgbotapi.NewEditMessageReplyMarkup(chatID, callback.Message.MessageID, newKeyboard)
		if _, err := h.bot.Send(edit); err != nil {
			log.Printf("Ошибка обновления клавиатуры: %v", err)
		}

	case buttons.PromotionsButton.Data:
		text = "🔥 Актуальные акции:"
	case buttons.ReviewsButton.Data:
		text = "⭐ Отзывы наших клиентов:"
	default:
		text = "Неизвестная кнопка"
	}

	msg := tgbotapi.NewMessage(chatID, text)
	h.bot.Send(msg)

	// Подтверждаем callback, чтобы Telegram не показывал "Загрузка..."
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := h.bot.Request(callbackConfig); err != nil {
		log.Printf("Ошибка отправки callback ответа: %v", err)
	}
}
