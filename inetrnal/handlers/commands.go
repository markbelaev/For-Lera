package handlers

import "gopkg.in/telebot.v4"

func handleStart(c telebot.Context) error {
	return c.Send("Привет")
}

func handleCompliments(c telebot.Context) error {
	compliments := []string{
		"Ты самая красивая",
		"Ты самая умная",
		"Ты самая сильная",
		"У тебя самые красивые глаза",
		"У тебя самые сочные губы",
		"Ты самая сексуальная",
		"Ты самая желанная",
		"Я теку от тебя",
		"Ты лучшая мама",
	}

	rand.Seed(time.Now().UnixNano())
	randomCompliments := compliments[rand.Intn(len(compliments))]

	return c.Send(randomCompliments)
}
