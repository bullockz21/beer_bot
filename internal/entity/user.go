package entity

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey"`
	TelegramID int64     `gorm:"uniqueIndex;not null"`
	Username   string    `gorm:"size:255"`
	FirstName  string    `gorm:"size:255;not null"`
	Language   string    `gorm:"size:2;default:'ru'"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
