package telegram

import (
	"context"
	"net/http"

	"github.com/bullockz21/beer_bot/internal/bot"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

// WebhookHandler godoc
// @Summary Webhook Endpoint
// @Description Receives incoming updates from Telegram
// @Tags Webhook
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/webhook [post]
func WebhookHandler(handler *bot.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// PingHandler возвращает pong для тестирования сервера.
// @Summary Ping Endpoint
// @Description Returns pong message for testing
// @Tags Example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
