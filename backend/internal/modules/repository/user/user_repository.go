// internal/repository/user_repository.go
package repository

import (
	"context"

	"github.com/bullockz21/beer_bot/internal/modules/domain"
	entityUser "github.com/bullockz21/beer_bot/internal/modules/entity/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Функция для преобразования доменной модели в entity.
func domainToEntity(u *domain.User) *entityUser.User {
	return &entityUser.User{
		TelegramID: u.TelegramID,
		Username:   u.Username,
		FirstName:  u.FirstName,
		Language:   u.Language,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// Функция для преобразования entity в доменную модель.
func entityToDomain(e *entityUser.User) (*domain.User, error) {
	u, err := domain.NewUser(e.TelegramID, e.Username, e.FirstName, e.Language)
	if err != nil {
		return nil, err
	}
	u.CreatedAt = e.CreatedAt
	u.UpdatedAt = e.UpdatedAt
	return u, nil
}

func (r *UserRepository) Save(ctx context.Context, u *domain.User) error {
	eUser := domainToEntity(u)
	return r.db.WithContext(ctx).
		Where(&entityUser.User{TelegramID: eUser.TelegramID}).
		Assign(eUser).
		FirstOrCreate(eUser).
		Error
}

func (r *UserRepository) FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	var eUser entityUser.User
	if err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&eUser).Error; err != nil {
		return nil, err
	}
	return entityToDomain(&eUser)
}
