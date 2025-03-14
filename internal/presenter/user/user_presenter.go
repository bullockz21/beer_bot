package presenter

import (
	"fmt"

	"github.com/bullockz21/beer_bot/internal/domain/buttons"
	dto "github.com/bullockz21/beer_bot/internal/dto/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

// presenter/user_presenter.go
func (p *UserPresenter) PresentUserSuccess(chatID int64, user *dto.UserResponse) error {
	// Создаем инлайн-клавиатуру из предопределенных кнопок
	keyboard := buttons.InlineKeyboard(
		buttons.MenuButton,
		buttons.PromotionsButton,
		buttons.ReviewsButton,
	)

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
		"🎉 Добро пожаловать, %s!\nВаш аккаунт успешно зарегистрирован!\n\nID: %d\nUsername: @%s",
		user.FirstName,
		user.ID,
		user.Username,
	))
	// Прикрепляем клавиатуру к сообщению
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "Markdown"

	_, err := p.bot.Send(msg)
	return err
}

func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msg := tgbotapi.NewMessage(chatID, "🚫 Ошибка: "+errorMsg)
	_, err := p.bot.Send(msg)
	return err
}
