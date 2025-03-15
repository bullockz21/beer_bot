package main

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/bullockz21/beer_bot/configs"
	botPkg "github.com/bullockz21/beer_bot/internal/bot"
	telegramController "github.com/bullockz21/beer_bot/internal/controller/telegram"
	dbpkg "github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
	userPresenterPkg "github.com/bullockz21/beer_bot/internal/presenter/user"
	userRepositoryPkg "github.com/bullockz21/beer_bot/internal/repository/user"
	userResourcePkg "github.com/bullockz21/beer_bot/internal/resource/user"
	userUsecasePkg "github.com/bullockz21/beer_bot/internal/usecase/user"
)

func main() {
	// Загрузка конфигурации
	cfg, err := configs.Load()
	if err != nil {
		logrus.WithField("module", "config").Fatalf("Config error: %v", err)
	}

	// Инициализация базы данных
	db, err := dbpkg.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer func() {
		if err := dbpkg.Close(db); err != nil {
			log.Printf("Ошибка закрытия базы данных: %v", err)
		}
	}()

	// Применяем миграции
	if err := migration.Run(db); err != nil {
		log.Fatalf("Миграция не удалась: %v", err)
	}

	// Создаем Telegram-бота (файл internal/bot/bot.go)
	bot, err := botPkg.NewBot(cfg)
	if err != nil {
		log.Fatalf("Bot init failed: %v", err)
	}

	// Инициализация бизнес-логики
	userRepo := userRepositoryPkg.NewUserRepository(db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(bot)
	userResource := userResourcePkg.NewUserResource()

	// Создаем обработчики команд и callback
	// Создаем обработчики
	startHandler := telegramController.NewStartHandler(userUC, userPresenter)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter, userResource)
	callbackHandler := telegramController.NewCallbackHandler(bot)
	handler := telegramController.NewHandler(bot, commandHandler, callbackHandler)

	// Запускаем получение обновлений (файл internal/bot/updates.go)
	ctx := context.Background()
	botPkg.ListenUpdates(ctx, bot, handler)
}
