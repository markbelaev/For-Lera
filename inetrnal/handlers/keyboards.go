package handlers

import "gopkg.in/telebot.v4"

func MenuKeyboard() *telebot.ReplyMarkup {
	menuKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnAslLoveMe := menuKeyboard.Text("Ты меня любишь?")
	btnTimeLove := menuKeyboard.Text("Сколько мы знакомы")

	menuKeyboard.Reply(
		menuKeyboard.Row(btnAslLoveMe, btnTimeLove),
	)

	return menuKeyboard
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
