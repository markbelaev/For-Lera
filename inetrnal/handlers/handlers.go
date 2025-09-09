package handlers

import "gopkg.in/telebot.v4"

func RegisterAll(b *telebot.Bot) {
	registerCommands(b)
}

func registerCommands(b *telebot.Bot) {
	b.Handle("/start", handleStart)
	b.Handle("/compliments", handleCompliments)
	b.Handle("/dreams", handleDreams)
	b.Handle("/motivations", handleMotivation)
}
