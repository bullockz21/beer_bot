package main

import (
	"log"

	"github.com/bullockz21/beer_bot/configs"
	_ "github.com/bullockz21/beer_bot/docs" // анонимный импорт генерированных swagger файлов
	botPkg "github.com/bullockz21/beer_bot/internal/bot"
	telegramController "github.com/bullockz21/beer_bot/internal/controller/telegram"
	dbpkg "github.com/bullockz21/beer_bot/internal/infrastructure/database"
	"github.com/bullockz21/beer_bot/internal/infrastructure/migration"
	userPresenterPkg "github.com/bullockz21/beer_bot/internal/presenter/user"
	userRepositoryPkg "github.com/bullockz21/beer_bot/internal/repository/user"
	"github.com/bullockz21/beer_bot/internal/router"
	userUsecasePkg "github.com/bullockz21/beer_bot/internal/usecase/user"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"     // Swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger UI handler
)

func main() {
	// Устанавливаем режим Gin (для продакшна обычно ReleaseMode)
	gin.SetMode(gin.ReleaseMode)

	// Загружаем конфигурацию
	cfg, err := configs.Load()
	if err != nil {
		logrus.WithField("module", "config").Fatalf("Config error: %v", err)
	}

	// Инициализируем базу данных
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

	// Настройка вебхука для Telegram (публичный URL должен быть корректным и включать версионный префикс)
	webhookURL := cfg.WebhookURL + "/api/v1/webhook"
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}
	if _, err = bot.Request(webhookConfig); err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	// Инициализируем зависимости для бизнес-логики
	userRepo := userRepositoryPkg.NewUserRepository(db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(bot)
	startHandler := telegramController.NewStartHandler(userUC, userPresenter, cfg)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(bot)
	handler := telegramController.NewHandler(bot, commandHandler, callbackHandler)

	// Настраиваем маршруты через модуль router
	r := router.SetupRoutes(handler)

	// Добавляем эндпоинт для Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера Gin на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
