package main

import (
	"log"

	"github.com/bullockz21/beer_bot/configs"
	"github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Printf("Database close error: %v", err)
		}
	}()

	if err := migration.Run(db); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	// Инициализация и запуск бота
	log.Println("Application started successfully")
}
