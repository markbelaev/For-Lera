package handlers

import "gopkg.in/telebot.v4"

func MenuKeyboard() *telebot.ReplyMarkup {
	menuKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnTestLoveKeyboard := menuKeyboard.Text("Ты меня любишь? 🤔")
	btnTimeLove := menuKeyboard.Text("Сколько знакомы ⏰")

	btnCompliments := menuKeyboard.Text("❤️")
	btnMotivations := menuKeyboard.Text("✨")
	btnDreams := menuKeyboard.Text("💫")

	menuKeyboard.Reply(
		menuKeyboard.Row(btnTestLoveKeyboard, btnTimeLove),
		menuKeyboard.Row(btnCompliments, btnMotivations, btnDreams),
	)

	return menuKeyboard
}

func TestLoveKeyboard() *telebot.ReplyMarkup {
	testLoveKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnYes := testLoveKeyboard.Text("ДА! 🙂‍↕️")
	btnNo := testLoveKeyboard.Text("не 🙂‍↔️")

	testLoveKeyboard.Reply(
		testLoveKeyboard.Row(btnYes, btnNo),
	)

	return testLoveKeyboard
}

func RemoveKeyboard() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	}
}
