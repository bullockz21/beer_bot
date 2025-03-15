// internal/usecase/user/user_usecase.go
package user

import (
	"context"
	"fmt"

	"github.com/bullockz21/beer_bot/internal/domain"
)

type UserUseCase struct {
	repo domain.Repository
}

func NewUserUseCase(repo domain.Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// CreateUser создает пользователя, используя доменную логику, и сохраняет его через репозиторий.
func (uc *UserUseCase) CreateUser(ctx context.Context, telegramID int64, username, firstName, language string) (*domain.User, error) {
	// Создаем доменного пользователя с 4 аргументами.
	user, err := domain.NewUser(telegramID, username, firstName, language)
	if err != nil {
		return nil, err
	}
	// Передаем контекст в вызов метода Save.
	if err := uc.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}
	return user, nil
}
