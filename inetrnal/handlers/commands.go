package handlers

import "gopkg.in/telebot.v4"

func handleStart(c telebot.Context) error {
	return c.Send("Привет")
}
