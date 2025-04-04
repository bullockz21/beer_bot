package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	TelegramToken string
}

func Load() (*Config, error) {
	// Загрузка переменных окружения из файла .env
	_ = godotenv.Load("../configs/.env")

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        port,
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}, nil
}
