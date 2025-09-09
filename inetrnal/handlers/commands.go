package handlers

import (
	"math/rand"

	"gopkg.in/telebot.v4"
)

func handleStart(c telebot.Context) error {
	return c.Send("Привет, котик ❤️. " +
		"\nНажимай кнопочки из меню и поднимай себе настроение: " +
		"\n\n/compliments — команда для получения комплимента (ты, кстати, сегодня классно выглядишь). " +
		"\n\n/motivations — команда для поднятия мотивации (твои мечты обязательно сбудутся). " +
		"\n\n/dreams — команда, где будут мои влажные мечты (у меня пубертат). " +
		"\n\nЯ буду обновлять бота, но пока все так.\n\nТвой Максимчик 💋")
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
		"Лучший тимейт ❤️💋",
		"У тебя идеальная фигура ❤️💋",
		"Ты очень талантливая ❤️💋",
		"Ты мое вдохновение ❤️💋",
		"Ты лучшая девушка на свете ❤️💋",
		"Ты делаешь этот мир лучше ❤️💋",
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

func handleMotivation(c telebot.Context) error {
	motivations := []string{
		"У тебя всё получится ✨",
		"Не сомневайся в себе - ты прекрасна ✨",
		"Ты сильнее, чем думаешь ✨",
		"Ты все сможешь ✨",
		"Твои мечты обязательно сбудутся ✨",
		"Продолжай двигаться вперед ✨",
		"Ты делаешь обычные дни особенными ✨",
	}

	randomIndex := rand.Intn(len(motivations))
	randomMotivation := motivations[randomIndex]

	return c.Send(randomMotivation)
}
