package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func help(update tgbotapi.Update) (string, interface{}) {
	/*
	TODO
	 */
	return ":help", nil
}

func createNewTrip(update tgbotapi.Update) (string, interface{}) {
	text := update.Message.Text

	pieces := strings.SplitN(text, " ", 2)

	if len(pieces) < 2 {
		return ":specify_trip", nil
	}

	// TODO give opportunity to change active trip
	trip, err := getCurrentTrip(update)

	if err == nil {
		return ":already_has_active_trip", nil
	}

	trip = Trip {
		Name: pieces[1],
		OwnerID: update.Message.From.ID,
		ChatID: update.Message.Chat.ID,
	}

	db.Create(&trip)

	trip.addMember(update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName)

	// TODO TODO TODO DETERMINE THE FUCK YOU DO WITH SUBSTITUTIONS
	return fmt.Sprintf("Successfully created new trip: %s", trip.Name), nil
}

func addMember(update tgbotapi.Update) (string, interface{}) {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error(), nil
	}

	username := update.Message.From.UserName
	userID := update.Message.From.ID

	var member TripMember

	record := db.Where("trip_id", trip.ID).Where("user_id", userID).First(&member)

	if record.Error == nil {
		return ":you_are_already_in", nil
	}

	trip.addMember(userID, username, update.Message.From.FirstName)

	return ":you_are_in", nil
}

func addDebt(update tgbotapi.Update) (string, interface{}) {
	pieces := strings.Split(update.Message.Text, " ")

	if len(pieces) < 3 {
		return ":add_usage", nil
	}

	amount, err := strconv.Atoi(pieces[1])

	if err != nil {
		return ":add_usage", nil
	}

	description := pieces[2]

	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error(), nil
	}

	expense, err := trip.addExpense(update.Message.From.ID, amount, description)

	if err != nil {
		return err.Error(), nil
	}

	expense.split(update.Message.From.ID)

	// TODO add more information
	return ":expense_created", nil
}

func getMembers(update tgbotapi.Update) (string, interface{}) {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return err.Error(), nil
	}

	var members []TripMember

	db.Where("trip_id", trip.ID).Find(&members)

	// TODO add member debts
	response := "\"" + trip.Name + "\" members:\n"

	for _, member := range members {
		response += fmt.Sprintf(" - %s (%s)\n", member.FirstName, member.Username)
	}

	return response, nil
	// TODO TODO TODO HERE TOO
}

func defaultAnswer(update tgbotapi.Update) (string, interface{}) {
	return ":default_answer", nil
}