package handlers

import (
	"For-Lera/inetrnal/data"
	"math/rand"

	"gopkg.in/telebot.v4"
)

func handleStart(c telebot.Context) error {
	return c.Send("Привет, котик ❤️. "+
		"\nНажимай кнопочки из меню и поднимай себе настроение: "+
		"\n\n/compliments — команда для получения комплимента (ты, кстати, сегодня классно выглядишь). "+
		"\n\n/motivations — команда для поднятия мотивации (твои мечты обязательно сбудутся). "+
		"\n\n/dreams — команда, где будут мои влажные мечты (у меня пубертат). "+
		"\n\nЯ буду обновлять бота, но пока все так.\n\nТвой Максимчик 💋",
		AskLoveKeyboard())
}

func handleCompliments(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Compliments))
	randomCompliment := data.Compliments[randomIndex]

	return c.Send(randomCompliment, AskLoveKeyboard())
}

func handleMotivation(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Motivations))
	randomMotivation := data.Motivations[randomIndex]

	return c.Send(randomMotivation, AskLoveKeyboard())
}

func handleDreams(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Dreams))
	randomDream := data.Dreams[randomIndex]

	return c.Send(randomDream, AskLoveKeyboard())
}

func handleLoveQuestion(c telebot.Context) error {
	return c.Send("Ты меня любишь, пупс? 🤔", LoveQuestionKeyboard())
}

func handleYes(c telebot.Context) error {
	return c.Send("Ты мне тоже нравишься ❤️", RemoveKeyboard())
}

func handleNo(c telebot.Context) error {
	return c.Send("динаху 😒", RemoveKeyboard())
}

func handleTimeLove(c telebot.Context) error {
	return c.Send("Если честно сам не знаю((( \n\n Потом посчитаю", RemoveKeyboard())
}
