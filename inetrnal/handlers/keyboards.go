package handlers

import "gopkg.in/telebot.v4"

func AskLoveKeyboard() *telebot.ReplyMarkup {
	askLoveKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnLoveMe := askLoveKeyboard.Text("Ты меня любишь?")

	askLoveKeyboard.Reply(
		askLoveKeyboard.Row(btnLoveMe),
	)

	return askLoveKeyboard
}

func LoveQuestionKeyboard() *telebot.ReplyMarkup {
	loveQuestionKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnYes := loveQuestionKeyboard.Text("ДА!")
	btnNo := loveQuestionKeyboard.Text("не")

	loveQuestionKeyboard.Reply(
		loveQuestionKeyboard.Row(btnYes, btnNo),
	)

	return loveQuestionKeyboard
}

func RemoveKeyboard() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	}
}
