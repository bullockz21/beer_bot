package migration

import (
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get sql.DB failed: %v", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	const migrationDir = "../migrations"

	if err := goose.Up(sqlDB, migrationDir); err != nil {
		return fmt.Errorf("миграция не удалась: %v", err)
	}

	log.Println("Миграции успешно применены")
	return nil
}
