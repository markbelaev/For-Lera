package handlers

import (
	"For-Lera/inetrnal/data"
	"math/rand"

	"gopkg.in/telebot.v4"
)

// Обработчик команды /start

func handleStart(c telebot.Context) error {
	return c.Send(data.StartMessage,
		MenuKeyboard())
}

// Обработчик команды /compliments

func handleCompliments(c telebot.Context) error {

	//Достаем из data список нужных сообщений и присылаем рандомное

	randomIndex := rand.Intn(len(data.Compliments))
	randomCompliment := data.Compliments[randomIndex]

	return c.Send(randomCompliment, MenuKeyboard())
}

// Обработчик команды /motivations

func handleMotivation(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Motivations))
	randomMotivation := data.Motivations[randomIndex]

	return c.Send(randomMotivation, MenuKeyboard())
}

// Обработчик команды /dreams

func handleDreams(c telebot.Context) error {
	randomIndex := rand.Intn(len(data.Dreams))
	randomDream := data.Dreams[randomIndex]

	return c.Send(randomDream, MenuKeyboard())
}

// Обработчик для проверки чувств

func handleLoveQuestion(c telebot.Context) error {
	return c.Send("Ты меня любишь, пупс? 🤔", TestLoveKeyboard())
}

// Обработчик если на тест ответили "да"

func handleYes(c telebot.Context) error {
	return c.Send("Ты мне тоже нравишься ❤️", MenuKeyboard())
}

// Обработчик есди на тест ответили "нет"

func handleNo(c telebot.Context) error {
	return c.Send("динаху 😒", MenuKeyboard())
}

// Обработчик на копку "сколько закомы"

func handleTimeLove(c telebot.Context) error {
	return c.Send("Если честно сам не знаю((( \n\n Потом посчитаю", MenuKeyboard())
}
