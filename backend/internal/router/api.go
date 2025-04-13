package router

import (
	"context"
	"net/http"

	"github.com/bullockz21/beer_bot/internal/controller/telegram"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

// SetupRoutes настраивает все маршруты Gin и возвращает роутер.
func SetupRoutes(handler *telegram.Handler) *gin.Engine {
	// Создаем стандартный роутер Gin
	router := gin.Default()

	// Создаем группу маршрутов с префиксом /api/v1
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/webhook", func(c *gin.Context) {
			logrus.Info("🔥 Вебхук вызван!")
			var update tgbotapi.Update
			if err := c.ShouldBindJSON(&update); err != nil {
				logrus.Errorf("❌ Ошибка разбора JSON: %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			logrus.Infof("✅ Update получен: %+v", update)
			go handler.ProcessUpdate(context.Background(), update)
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Дополнительный тестовый маршрут
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	return router
}
