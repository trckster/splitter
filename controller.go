package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
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
		OwnerID: update.Message.From.ID,
		ChatID: update.Message.Chat.ID,
	}

	db.Create(&trip)

	trip.addMember(update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName)

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

	trip.addMember(userID, username, update.Message.From.FirstName)

	return "Done, you're in!"
}

func addDebt(update tgbotapi.Update) string {
	pieces := strings.Split(update.Message.Text, " ")

	usage := "Usage: /add <sum> <description>"

	if len(pieces) < 3 {
		return usage
	}

	amount, err := strconv.Atoi(pieces[1])

	if err != nil {
		return usage
	}

	description := pieces[2]

	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error()
	}

	expense, err := trip.addExpense(update.Message.From.ID, amount, description)

	if err != nil {
		return err.Error()
	}

	expense.split(update.Message.From.ID)

	// TODO add more information
	return "Expense created"
}

func getMembers(update tgbotapi.Update) string {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error()
	}

	var members []TripMember

	db.Where("trip_id", trip.ID).Find(&members)

	// TODO add member debts
	response := "\"" + trip.Name + "\" members:\n"

	for _, member := range members {
		response += fmt.Sprintf(" - %s (%s)\n", member.FirstName, member.Username)
	}

	return response
}

func defaultAnswer(update tgbotapi.Update) string {
	return "Oh, man, I don't know what are you talking about!"
}