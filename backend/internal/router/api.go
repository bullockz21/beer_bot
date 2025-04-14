// @title Beer Bot API
// @version 1.0
// @description Backend API for Telegram Beer Bot
// @contact.name Кирилл
// @contact.email kirill@example.com
// @host localhost:8080
// @BasePath /api/v1
package router

import (
	"github.com/bullockz21/beer_bot/internal/controller/telegram"
	"github.com/gin-gonic/gin"
)

// SetupRoutes настраивает все маршруты Gin и возвращает роутер.
func SetupRoutes(handler *telegram.Handler) *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/webhook", telegram.WebhookHandler(handler))
		apiV1.GET("/ping", telegram.PingHandler)
	}

	return router
}
