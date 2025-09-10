package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Структура где хранится токен бота: TOKEN_BOT=...

type Config struct {
	TokenBot string
}

// Функция загрузки конфига

func Load() *Config {

	// Загружем .env файл где хранится токен бота: TOKEN_BOT=...

	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	// Достаем токен бота: TOKEN_BOT=...

	tokenBot := os.Getenv("TOKEN_BOT")
	if tokenBot == "" {
		slog.Error("Token is empty")
		os.Exit(1)
	}

	return &Config{
		TokenBot: tokenBot,
	}
}
