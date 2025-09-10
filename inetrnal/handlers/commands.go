package handlers

import (
	"For-Lera/inetrnal/data"
	"math/rand"

	"gopkg.in/telebot.v4"
)

func handleStart(c telebot.Context) error {
	return c.Send(data.StartMessage,
		MenuKeyboard())
}

func handleCompliments(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Compliments))
	randomCompliment := data.Compliments[randomIndex]

	return c.Send(randomCompliment, MenuKeyboard())
}

func handleMotivation(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Motivations))
	randomMotivation := data.Motivations[randomIndex]

	return c.Send(randomMotivation, MenuKeyboard())
}

func handleDreams(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Dreams))
	randomDream := data.Dreams[randomIndex]

	return c.Send(randomDream, MenuKeyboard())
}

func handleLoveQuestion(c telebot.Context) error {
	return c.Send("Ты меня любишь, пупс? 🤔", TestLoveKeyboard())
}

func handleYes(c telebot.Context) error {
	return c.Send("Ты мне тоже нравишься ❤️", MenuKeyboard())
}

func handleNo(c telebot.Context) error {
	return c.Send("динаху 😒", MenuKeyboard())
}

func handleTimeLove(c telebot.Context) error {
	return c.Send("Если честно сам не знаю((( \n\n Потом посчитаю", MenuKeyboard())
}
