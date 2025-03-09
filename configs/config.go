// configs/config.go
package configs

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	TelegramToken string
	AdminIDs      []int64
}

func Load() (*Config, error) {
	_ = godotenv.Load("../configs/.env") // Try load .env, ignore error

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	adminIDs := parseAdminIDs(os.Getenv("ADMIN_IDS"))

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        port,
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		AdminIDs:      adminIDs,
	}, nil
}

func parseAdminIDs(input string) []int64 {
	var ids []int64
	for _, s := range strings.Split(input, ",") {
		id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}
