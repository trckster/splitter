package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Fixed keyboards

var createTripKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("create-trip", "?"),
	),
)

var tripKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("", ""),
	),
)