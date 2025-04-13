package main

import (
	"log"

	"github.com/bullockz21/beer_bot/configs"
	botPkg "github.com/bullockz21/beer_bot/internal/bot"
	telegramController "github.com/bullockz21/beer_bot/internal/controller/telegram"
	dbpkg "github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
	userPresenterPkg "github.com/bullockz21/beer_bot/internal/presenter/user"
	userRepositoryPkg "github.com/bullockz21/beer_bot/internal/repository/user"
	userUsecasePkg "github.com/bullockz21/beer_bot/internal/usecase/user"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	// Импортируем наш модуль с роутами:
	"github.com/bullockz21/beer_bot/internal/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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

	// Инициализация Telegram-бота
	bot, err := botPkg.NewBot(cfg)
	if err != nil {
		log.Fatalf("Bot init failed: %v", err)
	}
	log.Printf("Бот запущен: %s", bot.Self.UserName)

	// Настройка вебхука
	webhookURL := cfg.WebhookURL + "/api/v1/webhook" // Убедитесь, что это правильный URL
	// замените на актуальный публичный HTTPS URL
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}
	if _, err = bot.Request(webhookConfig); err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	// Инициализация зависимостей для бизнес-логики
	userRepo := userRepositoryPkg.NewUserRepository(db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(bot)
	startHandler := telegramController.NewStartHandler(userUC, userPresenter, cfg)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(bot)
	handler := telegramController.NewHandler(bot, commandHandler, callbackHandler)

	// Используем нашу функцию для настройки маршрутов
	r := router.SetupRoutes(handler)

	// Запуск сервера Gin (на порту 8080)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
