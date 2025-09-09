package handlers

import (
	"math/rand"

	"gopkg.in/telebot.v4"
)

func handleStart(c telebot.Context) error {
	return c.Send("Привет")
}

func handleCompliments(c telebot.Context) error {
	compliments := []string{
		"Ты самая красивая ❤️💋",
		"Ты самая умная ❤️💋",
		"Ты самая сильная ❤️💋",
		"У тебя самые красивые глаза ❤️💋",
		"У тебя самые сочные губы ❤️💋",
		"Ты самая сексуальная ❤️💋",
		"Ты самая желанная ❤️💋",
		"Я теку от тебя ❤️💋",
		"Ты самая лучшая мама ❤️💋",
		"Я пьян от тебя ❤️💋",
		"У тебя самый сочный зад ❤️💋",
		"У тебя самая вкусная фигура ❤️💋",
		"Ты лучшая мама на свете ❤️💋",
		"У тебя самая очаровательная улыбка ❤️💋",
		"Ты прекрасна ❤️💋",
		"Ты прекрасна как рассвет ❤️💋",
	}

	randomIndex := rand.Intn(len(compliments))
	randomCompliment := compliments[randomIndex]

	return c.Send(randomCompliment)
}

func handleDreams(c telebot.Context) error {
	dreams := []string{
		"Хочу поцеловать тебя 💫",
		"Хочу тебя 💫",
		"Хочу стать твоим стулом 💫",
		"Хочу чтобы ты на меня села 💫",
		"Хочу тебе отлизать 💫",
		"Хочу твое милое фото 💫",
		"Хочу фото твоих ножек 💫",
		"Хочу фото твоей сочной попы 💫",
		"Хочу фото твоей груди 💫",
	}

	randomIndex := rand.Intn(len(dreams))
	randomDream := dreams[randomIndex]

	return c.Send(randomDream)
}
