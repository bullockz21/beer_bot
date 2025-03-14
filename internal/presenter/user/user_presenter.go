package user

import (
	"fmt"

	"github.com/bullockz21/beer_bot/internal/domain/buttons"
	dtoUser "github.com/bullockz21/beer_bot/internal/dto/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

func (p *UserPresenter) PresentUserSuccess(chatID int64, user *dtoUser.UserResponse) error {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(
		"🎉 Добро пожаловать, %s!\nВаш аккаунт успешно зарегистрирован!\n\nID: %d\nUsername: @%s",
		user.FirstName,
		user.ID,
		user.Username,
	))
	msg.ParseMode = "Markdown"
	_, err := p.bot.Send(msg)
	return err
}

func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msg := tgbotapi.NewMessage(chatID, "🚫 Ошибка: "+errorMsg)
	_, err := p.bot.Send(msg)
	return err
}

// PresentWelcomeMessage отправляет приветственное сообщение с инлайн-клавиатурой.
func (p *UserPresenter) PresentWelcomeMessage(chatID int64, firstName string) error {
	text := fmt.Sprintf("Привет, %s! Выберите, что вас интересует:", firstName)
	msg := tgbotapi.NewMessage(chatID, text)

	// Используем универсальную клавиатуру
	msg.ReplyMarkup = buttons.InlineKeyboard(buttons.MenuButton, buttons.PromotionsButton, buttons.ReviewsButton)

	_, err := p.bot.Send(msg)
	return err
}
