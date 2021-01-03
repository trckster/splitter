package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func help(update tgbotapi.Update) string {
	/*
	TODO
	 */
	return "Wanna some help? Go read documentation dude"
}

func createNewTrip(update tgbotapi.Update) string {
	text := update.Message.Text

	pieces := strings.SplitN(text, " ", 2)

	if len(pieces) < 2 {
		return "You should specify trip name.\n\nExample:\n/new Vacation in Germany"
	}

	var trip Trip

	record := db.Where("chat_id", update.Message.Chat.ID).First(&trip)

	// TODO give opportunity to change active trip
	if record.Error == nil {
		return "You already have an active trip in this chat"
	}

	trip = Trip {
		Name: pieces[1],
		OwnerId: update.Message.From.ID,
		ChatId: update.Message.Chat.ID,
	}

	db.Create(&trip)

	trip.addMember(update.Message.From.ID, update.Message.From.UserName)

	return "Successfully created new trip: " + trip.Name
}

func defaultAnswer(update tgbotapi.Update) string {
	return "Oh, man, I don't know what are you talking about!"
}