package usecase

import (
	"context"
	"fmt"

	"github.com/bullockz21/beer_bot/internal/dto"
	"github.com/bullockz21/beer_bot/internal/entity"
	"github.com/bullockz21/beer_bot/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) HandleStart(ctx context.Context, req *dto.UserCreateRequest) (*entity.User, error) {
	user := &entity.User{
		TelegramID: req.TelegramID,
		Username:   req.Username,
		FirstName:  req.FirstName,
		Language:   req.Language,
	}

	if err := uc.repo.CreateOrUpdate(ctx, user); err != nil {
		return nil, fmt.Errorf("user creation failed: %w", err)
	}

	return user, nil
}
