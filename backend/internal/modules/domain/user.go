package domain

import (
	"context"
	"errors"
	"time"
)

type UserID int64

type User struct {
	ID         UserID
	TelegramID int64
	Username   string
	FirstName  string
	Language   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUser(telegramID int64, username, firstName, language string) (*User, error) {
	if telegramID <= 0 {
		return nil, errors.New("некорректный telegramID")
	}
	if username == "" {
		return nil, errors.New("username не может быть пустым")
	}
	if firstName == "" {
		return nil, errors.New("firstName не может быть пустым")
	}
	if language == "" {
		language = "ru"
	}
	now := time.Now()
	return &User{
		TelegramID: telegramID,
		Username:   username,
		FirstName:  firstName,
		Language:   language,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

func (u *User) UpdateUsername(newUsername string) error {
	if newUsername == "" {
		return errors.New("username не может быть пустым")
	}
	u.Username = newUsername
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) IsValid() bool {
	return u.TelegramID > 0 && u.Username != "" && u.FirstName != ""
}

// Repository теперь принимает context и возвращает доменный объект
type Repository interface {
	Save(ctx context.Context, user *User) error
	FindByTelegramID(ctx context.Context, telegramID int64) (*User, error)
}
