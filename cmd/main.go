package main

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/bullockz21/beer_bot/configs"
	"github.com/bullockz21/beer_bot/internal/bot"
	"github.com/bullockz21/beer_bot/internal/controller/telegram"
	"github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
	presenter "github.com/bullockz21/beer_bot/internal/presenter/user"
	repository "github.com/bullockz21/beer_bot/internal/repository/user"
	resource "github.com/bullockz21/beer_bot/internal/resource/user"
	usecase "github.com/bullockz21/beer_bot/internal/usecase/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "config",
		}).Fatalf("Config error: %v", err)
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

	// Создание бота ДО инициализации зависимостей
	bot, err := bot.NewBot(cfg)
	if err != nil {
		log.Fatal("Bot init failed: %v".err)
	}

	// Инициализация слоев приложения
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	userPresenter := presenter.NewUserPresenter(bot) // Теперь bot доступен
	userResource := resource.NewUserResource()

	// Инициализация обработчиков с зависимостями
	handler := telegram.NewHandler(
		bot,
		userUC,
		userPresenter,
		userResource,
	)

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

	log.Println("Приложение успешно запущено")
}
