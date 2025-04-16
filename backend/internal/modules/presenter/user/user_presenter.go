package user

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

// PresentError отправляет сообщение об ошибке.
func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msgText := fmt.Sprintf("🚫 Ошибка: %s", errorMsg)
	msg := tgbotapi.NewMessage(chatID, msgText)
	_, err := p.bot.Send(msg)
	return err
}

// PresentWelcomeMessage отправляет приветственное сообщение с Web App кнопкой.
// miniAppURL — это HTTPS URL мини‑приложения.
func (p *UserPresenter) PresentWelcomeMessage(chatID int64, firstName, miniAppURL string) error {
	text := fmt.Sprintf("Привет, %s! На связи служба доставки \"Рыба и Рис\".\nНажмите кнопку ниже, чтобы сделать заказ.", firstName)
	webAppButton := tgbotapi.NewInlineKeyboardButtonWebApp("Сделать заказ", tgbotapi.WebAppInfo{
		URL: miniAppURL,
	})
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(webAppButton))
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	_, err := p.bot.Send(msg)
	return err

}
