package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Fixed keyboards

var createTripKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("create-trip", "ask-name"),
	),
)

// TODO add handlers for these buttons
var tripKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("my-debts", "??"),
		tgbotapi.NewInlineKeyboardButtonData("members", "??"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("new-expense", "??"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("history", "??"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("my-expenses", "??"),
	),
)

func defineKeyboardsTexts() {
	keyboardsTexts = make(map[string]map[string]string)

	keyboardsTexts["en"] = make(map[string]string)
	keyboardsTexts["ru"] = make(map[string]string)

	// English

	keyboardsTexts["en"]["create-trip"] = "Create new trip"

	keyboardsTexts["en"]["my-debts"] = "My debts"
	keyboardsTexts["en"]["members"] = "Members"
	keyboardsTexts["en"]["new-expense"] = "New expense"
	keyboardsTexts["en"]["history"] = "History"
	keyboardsTexts["en"]["my-expenses"] = "My expenses"

	keyboardsTexts["en"]["dummy-button"] = "Dummy button"

	// Russian

	keyboardsTexts["ru"]["create-trip"] = "Создать поездку"

	keyboardsTexts["ru"]["my-debts"] = "Мои долги"
	keyboardsTexts["ru"]["members"] = "Участники"
	keyboardsTexts["ru"]["new-expense"] = "Новая трата"
	keyboardsTexts["ru"]["history"] = "История"
	keyboardsTexts["ru"]["my-expenses"] = "Мои траты"

	keyboardsTexts["ru"]["dummy-button"] = "Глупая кнопочка"
}
