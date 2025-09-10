package main

import (
	"For-Lera/inetrnal/config"
	"For-Lera/inetrnal/handlers"
	"log/slog"
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func main() {

	// Инициализация текстового логера

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// Загрузка конфиг файла

	cfg := config.Load()

	// Настройка бота

	pref := telebot.Settings{
		Token: cfg.TokenBot,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	// Создание бота

	bot, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("Failed to create bot", "error", err)
	}

	// Регистрация обработчиков

	handlers.RegisterAll(bot)

	// Запуск бота

	slog.Info("Bot started")
	bot.Start()
}
