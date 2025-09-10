package handlers

import (
	"For-Lera/inetrnal/data"
	"log/slog"
	"math/rand"

	"gopkg.in/telebot.v4"
)

// Обработчик команды /start

func handleStart(c telebot.Context) error {

	// Логируем успешную обработку

	slog.Info("Успешная обработка комманды /start!")

	// Возвращаем ответ

	return c.Send(data.StartMessage,
		MenuKeyboard())
}

// Обработчик команды /compliments

func handleCompliments(c telebot.Context) error {
	slog.Info("Успешная обработка комманды /compliments!")

	//Достаем из data список нужных сообщений и присылаем рандомное

	randomIndex := rand.Intn(len(data.Compliments))
	randomCompliment := data.Compliments[randomIndex]

	return c.Send(randomCompliment, MenuKeyboard())
}

// Обработчик команды /motivations

func handleMotivation(c telebot.Context) error {
	slog.Info("Успешная обработка комманды /motivations!")

	randomIndex := rand.Intn(len(data.Motivations))
	randomMotivation := data.Motivations[randomIndex]

	return c.Send(randomMotivation, MenuKeyboard())
}

// Обработчик команды /dreams

func handleDreams(c telebot.Context) error {
	slog.Info("Успешная обработка комманды /dreams!")

	randomIndex := rand.Intn(len(data.Dreams))
	randomDream := data.Dreams[randomIndex]

	return c.Send(randomDream, MenuKeyboard())
}

// Обработчик грустных сообщений

func handleSad(c telebot.Context) error {
	slog.Info("Успешня обработка комманды /sad!")

	randomIndex := rand.Intn(len(data.Sad))
	randomSad := data.Sad[randomIndex]

	return c.Send(randomSad, MenuKeyboard())
}

// Обработчик для проверки чувств

func handleLoveQuestion(c telebot.Context) error {
	slog.Info("Успешная обработка кнопки test love!")

	return c.Send("Ты меня любишь, пупс? 🤔", TestLoveKeyboard())
}

// Обработчик если на тест ответили "да"

func handleYes(c telebot.Context) error {
	slog.Info("Успешная обработка кнопки yes!")

	return c.Send("Ты мне тоже нравишься ❤️", MenuKeyboard())
}

// Обработчик есди на тест ответили "нет"

func handleNo(c telebot.Context) error {
	slog.Info("Успешная обработка кнопки no!")

	return c.Send("динаху 😒", MenuKeyboard())
}

// Обработчик на копку "сколько закомы"

func handleTimeLove(c telebot.Context) error {
	slog.Info("Успешная обработка кнопки time love!")

	return c.Send("Если честно сам не знаю((( \n\n Потом посчитаю", MenuKeyboard())
}

func handleAnyMessages(c telebot.Context) error {
	userMessage := c.Text()

	slog.Info("Неизвестное сообщение: ", userMessage)

	return c.Send("Не понимаю о чем ты 🤷‍♂️", MenuKeyboard())
}
