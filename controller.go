package main

import (
	"fmt"
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

	// TODO give opportunity to change active trip
	trip, err := getCurrentTrip(update)

	if err == nil {
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

func addMember(update tgbotapi.Update) string {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error()
	}

	username := update.Message.From.UserName
	userID := update.Message.From.ID

	var member TripMember

	record := db.Where("trip_id", trip.ID).Where("user_id", userID).First(&member)

	if record.Error == nil {
		return "You're already in the trip!"
	}

	trip.addMember(userID, username)

	return "Done, you're in!"
}

func addDebt(update tgbotapi.Update) string {
	// TODO
	return "Work in progress"
}

func getMembers(update tgbotapi.Update) string {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error()
	}

	var members []TripMember

	db.Where("trip_id", trip.ID).Find(&members)

	// TODO add member debts
	response := "Trip members:\n"

	for _, member := range members {
		response += fmt.Sprintf(" - %s\n", member.Username)
	}

	return response
}

func defaultAnswer(update tgbotapi.Update) string {
	return "Oh, man, I don't know what are you talking about!"
}