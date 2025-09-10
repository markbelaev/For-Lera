package handlers

import "gopkg.in/telebot.v4"

// Функция регистрации всх обработчиков

func RegisterAll(b *telebot.Bot) {
	registerCommands(b)
	registerAnyMessages(b)
}

// Хендлеры

func registerCommands(b *telebot.Bot) {

	// Обработчики комманд

	b.Handle("/start", handleStart)
	b.Handle("/compliments", handleCompliments)
	b.Handle("/motivations", handleMotivation)
	b.Handle("/dreams", handleDreams)
	b.Handle("/sad", handleSad)

	// Обработчики для алиасов комманд

	b.Handle("❤️", handleCompliments)
	b.Handle("✨", handleMotivation)
	b.Handle("💫", handleDreams)
	b.Handle("😢", handleSad)

	// Обработчики для love test

	b.Handle("Ты меня любишь? 🤔", handleLoveQuestion)
	b.Handle("ДА! 🙂‍↕️", handleYes)
	b.Handle("не 🙂‍↔️", handleNo)

	// Остальные обработчики

	b.Handle("Сколько знакомы ⏰", handleTimeLove)
}

// Обработчик неизвестных текстовых сообщений

func registerAnyMessages(b *telebot.Bot) {
	b.Handle(telebot.OnText, handleAnyMessages)
}
