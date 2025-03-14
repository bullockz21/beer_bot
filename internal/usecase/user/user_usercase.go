package user

import (
	"context"
	"fmt"

	dtoUser "github.com/bullockz21/beer_bot/internal/dto/user"
	entityUser "github.com/bullockz21/beer_bot/internal/entity/user"
	repositoryUser "github.com/bullockz21/beer_bot/internal/repository/user"
)

type UserUseCase struct {
	repo repositoryUser.UserRepository
}

func NewUserUseCase(repo repositoryUser.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) HandleStart(ctx context.Context, req *dtoUser.UserCreateRequest) (*entityUser.User, error) {
	user := &entityUser.User{
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
