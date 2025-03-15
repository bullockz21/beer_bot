package domain

import (
	"errors"
	"time"
)

// UserID - тип для идентификатора пользователя
type UserID int64

// User - доменная сущность пользователя, содержащая бизнес-данные и логику.
type User struct {
	ID         UserID    // идентификатор пользователя (может задаваться базой данных)
	TelegramID int64     // уникальный ID из Telegram
	Username   string    // имя пользователя (логин)
	FirstName  string    // имя (реальное)
	CreatedAt  time.Time // время создания записи
	UpdatedAt  time.Time // время последнего обновления
}

// NewUser создаёт нового пользователя и проверяет базовые бизнес-правила.
func NewUser(telegramID int64, username, firstName string) (*User, error) {
	if telegramID <= 0 {
		return nil, errors.New("некорректный telegramID")
	}
	if username == "" {
		return nil, errors.New("username не может быть пустым")
	}
	now := time.Now()
	return &User{
		TelegramID: telegramID,
		Username:   username,
		FirstName:  firstName,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

// UpdateUsername обновляет имя пользователя с учётом бизнес-логики.
// Здесь можно добавить дополнительные проверки, например, по формату имени.
func (u *User) UpdateUsername(newUsername string) error {
	if newUsername == "" {
		return errors.New("username не может быть пустым")
	}
	u.Username = newUsername
	u.UpdatedAt = time.Now()
	return nil
}

// IsValid проверяет, что данные пользователя соответствуют базовым бизнес-правилам.
func (u *User) IsValid() bool {
	return u.TelegramID > 0 && u.Username != ""
}

// Repository описывает интерфейс для работы с данными пользователей.
// Реализацию этого интерфейса следует поместить в инфраструктурный слой.
type Repository interface {
	Save(user *User) error
	FindByTelegramID(telegramID int64) (*User, error)
}
