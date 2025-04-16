package main

import (
	"log"

	"github.com/bullockz21/beer_bot/configs"
	_ "github.com/bullockz21/beer_bot/docs"
	botPkg "github.com/bullockz21/beer_bot/internal/bot"
	telegramController "github.com/bullockz21/beer_bot/internal/bot"
	dbpkg "github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
	userPresenterPkg "github.com/bullockz21/beer_bot/internal/modules/presenter/user"
	userRepositoryPkg "github.com/bullockz21/beer_bot/internal/modules/repository/user"
	userUsecasePkg "github.com/bullockz21/beer_bot/internal/modules/usecase/user"
	"github.com/bullockz21/beer_bot/internal/router/v1"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Устанавливаем режим Gin (для продакшна обычно ReleaseMode)
	gin.SetMode(gin.ReleaseMode)

	cfg, err := configs.Load()
	if err != nil {
		logrus.WithField("module", "config").Fatalf("Config error: %v", err)
	}

	db, err := dbpkg.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer func() {
		if err := dbpkg.Close(db); err != nil {
			log.Printf("Ошибка закрытия базы данных: %v", err)
		}
	}()

	if err := migration.Run(db); err != nil {
		log.Fatalf("Миграция не удалась: %v", err)
	}

	bot, err := botPkg.NewBot(cfg)
	if err != nil {
		log.Fatalf("Bot init failed: %v", err)
	}
	log.Printf("Бот запущен: %s", bot.Self.UserName)

	webhookURL := cfg.WebhookURL + "/api/v1/webhook"
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}
	if _, err = bot.Request(webhookConfig); err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	userRepo := userRepositoryPkg.NewUserRepository(db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(bot)
	startHandler := telegramController.NewStartHandler(userUC, userPresenter, cfg)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(bot)
	handler := telegramController.NewHandler(bot, commandHandler, callbackHandler)

	r := router.SetupRoutes(handler)

	// Добавляем эндпоинт для Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера Gin на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
