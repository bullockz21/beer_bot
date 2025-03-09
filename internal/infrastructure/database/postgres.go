// internal/infrastructure/database/postgres.go
package database

import (
	"fmt"
	"log"

	"github.com/bullockz21/beer_bot/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подлкюченрия: %v", err)
	}

	log.Println("Успешное подключение к базе данных")
	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get raw DB error: %v", err)
	}
	return sqlDB.Close()
}
