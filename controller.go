package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func help(update tgbotapi.Update) string {
	return "Wanna some help? Go read documentation dude"
}

func createNewTrip(update tgbotapi.Update) string {
	text := update.Message.Text

	pieces := strings.SplitN(text, " ", 2)

	trip := Trip {
		Name: pieces[1],
		OwnerId: update.Message.From.ID,
		ChatId: update.Message.Chat.ID,
	}

	db.Create(&trip)

	return "Successfully created new trip: " + trip.Name
}

func defaultAnswer(update tgbotapi.Update) string {
	return "Oh, man, I don't know what are you talking about!"
}