package user

import (
	"context"

	entityUser "github.com/bullockz21/beer_bot/internal/entity/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateOrUpdate(ctx context.Context, user *entityUser.User) error
	FindByTelegramID(ctx context.Context, telegramID int64) (*entityUser.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateOrUpdate(ctx context.Context, user *entityUser.User) error {
	return r.db.WithContext(ctx).
		Where(entityUser.User{TelegramID: user.TelegramID}).
		Assign(user).
		FirstOrCreate(user).
		Error
}

func (r *userRepo) FindByTelegramID(ctx context.Context, telegramID int64) (*entityUser.User, error) {
	var user entityUser.User
	err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&user).Error
	return &user, err
}
