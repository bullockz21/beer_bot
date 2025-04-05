package main

import (
	"context"
	"log"
	"net/http"

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

	// Настройка webhook
	webhookURL := "https://5200-62-210-88-22.ngrok-free.app" // замените на ваш HTTPS URL и путь, например, /webhook
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatalf("Ошибка создания вебхука: %v", err)
	}
	_, err = bot.Request(webhookConfig)
	if err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	// Инициализация зависимостей для бизнес-логики
	userRepo := userRepositoryPkg.NewUserRepository(db)
	userUC := userUsecasePkg.NewUserUseCase(userRepo)
	userPresenter := userPresenterPkg.NewUserPresenter(bot)

	startHandler := telegramController.NewStartHandler(userUC, userPresenter)
	commandHandler := telegramController.NewCommandHandler(startHandler, userPresenter)
	callbackHandler := telegramController.NewCallbackHandler(bot)
	handler := telegramController.NewHandler(bot, commandHandler, callbackHandler)

	// Создаем Gin-роутер для обработки вебхука
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"}) // Доверять localhost и ngrok
	// Маршрут для вебхука. Telegram будет слать POST-запросы сюда.
	router.POST("/webhook", func(c *gin.Context) {
		log.Println("🔥 Вебхук вызван!") // ← добавь вот это

		var update tgbotapi.Update
		if err := c.ShouldBindJSON(&update); err != nil {
			log.Printf("❌ Ошибка разбора JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("✅ Update получен: %+v", update) // ← и вот это
		go handler.ProcessUpdate(context.Background(), update)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Запуск сервера Gin с поддержкой HTTPS.
	// Если у вас есть сертификаты, используйте RunTLS. Для разработки можно использовать ngrok.
	// Для dev режима — запускаем просто HTTP-сервер
	if err := router.RunTLS(":8443", "server.crt", "server.key"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}
