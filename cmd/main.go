package main

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/bullockz21/beer_bot/configs"
	"github.com/bullockz21/beer_bot/internal/controller/telegram"
	"github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"

	"github.com/bullockz21/beer_bot/internal/repository"
	"github.com/bullockz21/beer_bot/internal/usecase"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
	// Инициализация БД
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Printf("Ошибка закрытия базы данных: %v", err)
		}
	}()

	if err := migration.Run(db); err != nil {
		log.Fatalf("Миграция не удалась: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)

	// Создание бота
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Bot init failed: %v", err)
	}

	// Инициализация обработчиков
	handler := telegram.NewHandler(bot, userUC)

	// Настройка обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Обработка сообщений
	ctx := context.Background()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start":
			handler.HandleStart(ctx, update)
		}
	}

	// Инициализация и запуск бота
	log.Println("Приложение успешно запущено")
}
