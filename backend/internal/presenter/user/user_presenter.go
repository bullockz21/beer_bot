package user

import (
	"encoding/json"
	"fmt"

	"github.com/bullockz21/beer_bot/internal/presenter/customkeyboard"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// RawReplyMarkup — тип для передачи "сырых" JSON данных в поле ReplyMarkup.
type RawReplyMarkup struct {
	Data string
}

// MarshalJSON позволяет RawReplyMarkup удовлетворять интерфейсу json.Marshaler.
func (r RawReplyMarkup) MarshalJSON() ([]byte, error) {
	return []byte(r.Data), nil
}

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

// NewUserPresenter создаёт нового презентера для отправки сообщений.
func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

// PresentError отправляет сообщение об ошибке.
func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msg := tgbotapi.NewMessage(chatID, "🚫 Ошибка: "+errorMsg)
	_, err := p.bot.Send(msg)
	return err
}

// PresentWelcomeMessage отправляет приветственное сообщение с кнопкой для открытия мини‑аппа.
// miniAppURL — это HTTPS URL вашего мини‑приложения.
func (p *UserPresenter) PresentWelcomeMessage(chatID int64, firstName, miniAppURL string) error {
	text := fmt.Sprintf("Привет, %s! На связи служба доставки Рыба и Рис\n\nНажмите кнопку ниже, чтобы сделать заказ и посмотреть меню.\n\nРежим работы: 10:00-23:00\n\nКонтакты:\n📍 Адрес: Стройкерамика, ул.Березовая 35\n🙎‍♂️ По вопросам: @max888tr\n📞 89272013588", firstName)

	// Создаем клавиатуру с помощью нашей функции из pkg/customkeyboard.
	kb := customkeyboard.NewWebAppKeyboard("Сделать заказ", miniAppURL)
	keyboardJSON, err := json.Marshal(kb)
	if err != nil {
		return err
	}

	// Оборачиваем JSON в RawReplyMarkup для передачи в поле ReplyMarkup.
	rawMarkup := RawReplyMarkup{Data: string(keyboardJSON)}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = rawMarkup

	_, err = p.bot.Send(msg)
	return err
}
