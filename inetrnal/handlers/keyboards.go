package handlers

import "gopkg.in/telebot.v4"

// Клавиатура главного меню

func MenuKeyboard() *telebot.ReplyMarkup {

	// Инициализируем reply клавиатуру

	menuKeyboard := &telebot.ReplyMarkup{
		ResizeKeyboard: true,
	}

	// Инициализируем кнопки клавиатуры

	btnTestLoveKeyboard := menuKeyboard.Text("Ты меня любишь? 🤔")
	btnTimeLove := menuKeyboard.Text("Сколько знакомы ⏰")

	btnCompliments := menuKeyboard.Text("❤️")
	btnMotivations := menuKeyboard.Text("✨")
	btnDreams := menuKeyboard.Text("💫")
	btnSad := menuKeyboard.Text("😢")

	// Ставим кнопки в ряд

	menuKeyboard.Reply(
		menuKeyboard.Row(btnTestLoveKeyboard, btnTimeLove),
		menuKeyboard.Row(btnCompliments, btnMotivations, btnDreams, btnSad),
	)

	// Возвращаем клавиатуру

	return menuKeyboard
}

// Клавматура для проверки чувств

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

// Функция для закрытия reply клавиатуры

func RemoveKeyboard() *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{
		RemoveKeyboard: true,
	}
}
