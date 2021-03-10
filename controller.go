package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func help(update tgbotapi.Update) Answer {
	/*
		TODO
	*/
	return Answer{Signature: ":help"}
}

func createNewTrip(update tgbotapi.Update) Answer {
	text := update.Message.Text

	pieces := strings.SplitN(text, " ", 2)

	if len(pieces) < 2 {
		return Answer{Signature: ":specify_trip"}
	}

	// TODO give opportunity to change active trip
	trip, err := getCurrentTrip(update)

	if err == nil {
		return Answer{Signature: ":already_has_active_trip"}
	}

	trip = Trip{
		Name:    pieces[1],
		Language: update.Message.From.LanguageCode,
		OwnerID: update.Message.From.ID,
		ChatID:  update.Message.Chat.ID,
	}

	db.Create(&trip)

	trip.addMember(update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName)

	return Answer{
		Signature:  ":new_trip",
		Parameters: []string{trip.Name},
	}
}

func addMember(update tgbotapi.Update) Answer {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	username := update.Message.From.UserName
	userID := update.Message.From.ID

	var member TripMember

	record := db.Where("trip_id", trip.ID).Where("user_id", userID).First(&member)

	if record.Error == nil {
		return Answer{Signature: ":you_are_already_in"}
	}

	trip.addMember(userID, username, update.Message.From.FirstName)

	return Answer{Signature: ":you_are_in"}
}

func addDebt(update tgbotapi.Update) Answer {
	pieces := strings.Split(update.Message.Text, " ")

	if len(pieces) < 3 {
		return Answer{Signature: ":add_usage"}
	}

	amount, err := strconv.Atoi(pieces[1])

	if err != nil {
		return Answer{Signature: ":add_usage"}
	}

	description := pieces[2]

	trip, err := getCurrentTrip(update)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	expense, err := trip.addExpense(update.Message.From.ID, amount, description)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	expense.split(update.Message.From.ID)

	// TODO add more information
	return Answer{Signature: ":expense_created"}
}

func getMembers(update tgbotapi.Update) Answer {
	trip, err := getCurrentTrip(update)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	var members []TripMember

	db.Where("trip_id", trip.ID).Find(&members)

	// TODO add member debts
	response := "\"" + trip.Name + "\" members:\n"

	for _, member := range members {
		response += fmt.Sprintf(" - %s (%s)\n", member.FirstName, member.Username)
	}

	return Answer{
		Signature:  ":hack",
		Parameters: []string{response},
	}
}

func defaultAnswer(update tgbotapi.Update) Answer {
	return Answer{Signature: ":default_answer"}
}
