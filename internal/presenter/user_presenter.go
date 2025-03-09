// internal/presenter/user_presenter.go
package presenter

import (
	"fmt"

	"github.com/bullockz21/beer_bot/internal/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserPresenter struct {
	bot *tgbotapi.BotAPI
}

func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{bot: bot}
}

func (p *UserPresenter) PresentUserSuccess(chatID int64, user *dto.UserResponse) error {
	text := fmt.Sprintf(
		"👋 Привет, %s!\nID: %d\nUsername: @%s",
		user.FirstName,
		user.ID,
		user.Username,
	)
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := p.bot.Send(msg)
	return err
}

func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msg := tgbotapi.NewMessage(chatID, "🚫 Ошибка: "+errorMsg)
	_, err := p.bot.Send(msg)
	return err
}
