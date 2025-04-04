package user

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

// NewUserPresenter создаёт нового презентера для отправки сообщений
func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

// PresentError отправляет сообщение об ошибке
func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msg := tgbotapi.NewMessage(chatID, "🚫 Ошибка: "+errorMsg)
	_, err := p.bot.Send(msg)
	return err
}

// PresentWelcomeMessage отправляет приветственное сообщение с кнопкой для открытия мини‑аппа.
// miniAppURL — это HTTPS URL вашего мини‑приложения.
func (p *UserPresenter) PresentWelcomeMessage(chatID int64, firstName, miniAppURL string) error {
	text := fmt.Sprintf("Привет, %s! На связи служба доставки Рыба и Рис\n\nНажмите кнопку ниже, чтобы сделать заказ и посмотреть меню.\n\nРежим работы: 10:00-23:00\n\nКонтакты:\n📍 Адрес: Стройкерамика, ул.Березовая 35\n🙎‍♂️ По вопросам: @max888tr\n📞 89272013588", firstName)

	// Создаем кнопку Web App, которая откроет мини‑апп по заданному URL.
	webAppButton := tgbotapi.NewInlineKeyboardButtonWebApp("Сделать заказ", tgbotapi.WebAppInfo{
		URL: miniAppURL,
	})

	// Можно добавить дополнительные кнопки, если нужно. Здесь мы создаем одну строку с Web App кнопкой.
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(webAppButton),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard

	_, err := p.bot.Send(msg)
	return err
}
