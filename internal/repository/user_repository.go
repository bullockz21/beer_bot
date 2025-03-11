package repository

import (
	"context"

	"github.com/bullockz21/beer_bot/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateOrUpdate(ctx context.Context, user *entity.User) error
	FindByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository { //конструктор
	return &userRepo{db: db}
}

func (r *userRepo) CreateOrUpdate(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Where(entity.User{TelegramID: user.TelegramID}).
		Assign(user).
		FirstOrCreate(user).
		Error
}

func (r *userRepo) FindByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&user).Error
	return &user, err
}
