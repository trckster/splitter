package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func askName(update tgbotapi.Update) Answer {
	initState(update.CallbackQuery.From, update.CallbackQuery.Message.Chat, "CreateTrip")

	return Answer{Signature: "input-trip-name"}
}
