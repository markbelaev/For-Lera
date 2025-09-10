package handlers

import "gopkg.in/telebot.v4"

// Функция регистрации всх обработчиков

func RegisterAll(b *telebot.Bot) {
	registerCommands(b)
}

// Хендлеры

func registerCommands(b *telebot.Bot) {
	b.Handle("/start", handleStart)
	b.Handle("/compliments", handleCompliments)
	b.Handle("/motivations", handleMotivation)
	b.Handle("/dreams", handleDreams)

	b.Handle("❤️", handleCompliments)
	b.Handle("✨", handleMotivation)
	b.Handle("💫", handleDreams)

	b.Handle("Ты меня любишь? 🤔", handleLoveQuestion)
	b.Handle("ДА! 🙂‍↕️", handleYes)
	b.Handle("не 🙂‍↔️", handleNo)

	b.Handle("Сколько знакомы ⏰", handleTimeLove)
}
