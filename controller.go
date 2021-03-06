package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func start(update tgbotapi.Update) Answer {
	trip, err := getCurrentTrip(update.Message)

	if err == nil {
		return Answer{
			// TODO change this answer
			Signature:  "hack",
			Keyboard:   tripKeyboard,
			Parameters: map[string]string{":hack": "You already have a trip: *" + trip.Name + "*"},
			ParseMode:  ParseModeMarkdown,
		}
	}

	return Answer{Signature: "create-first-trip", Keyboard: createTripKeyboard}
}

func saveName(update tgbotapi.Update, state *FSM) Answer {
	text := update.Message.Text

	// TODO give opportunity to change active trip
	trip, err := getCurrentTrip(update.Message)

	if err == nil {
		state.next()
		return Answer{Signature: "already-has-active-trip"}
	}

	trip = Trip{
		Name:     text,
		Language: update.Message.From.LanguageCode,
		OwnerID:  update.Message.From.ID,
		ChatID:   update.Message.Chat.ID,
	}

	db.Create(&trip)

	state.next()

	trip.addMember(update.Message.From.ID, update.Message.From.UserName, update.Message.From.FirstName)

	return Answer{
		Signature:  "new-trip",
		Parameters: map[string]string{":trip_name": trip.Name},
	}
}

func addMember(update tgbotapi.Update) Answer {
	trip, err := getCurrentTrip(update.Message)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	username := update.Message.From.UserName
	userID := update.Message.From.ID

	var member TripMember

	record := db.Where("trip_id", trip.ID).Where("user_id", userID).First(&member)

	if record.Error == nil {
		return Answer{Signature: "you-are-already-in"}
	}

	trip.addMember(userID, username, update.Message.From.FirstName)

	return Answer{Signature: "you-are-in"}
}

func addDebt(update tgbotapi.Update) Answer {
	pieces := strings.Split(update.Message.Text, " ")

	if len(pieces) < 3 {
		return Answer{Signature: "add-usage"}
	}

	amount, err := strconv.Atoi(pieces[1])

	if err != nil {
		return Answer{Signature: "add-usage"}
	}

	description := pieces[2]

	trip, err := getCurrentTrip(update.Message)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	expense, err := trip.addExpense(update.Message.From.ID, amount, description)

	if err != nil {
		return Answer{Signature: err.Error()}
	}

	expense.split(update.Message.From.ID)

	// TODO add more information
	return Answer{Signature: "expense-created"}
}

func getMembers(update tgbotapi.Update) Answer {
	trip, err := getCurrentTrip(update.Message)

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
		Signature:  "hack",
		Parameters: map[string]string{":hack": response},
	}
}

func defaultAnswer(update tgbotapi.Update) Answer {
	return Answer{Signature: "default-answer"}
}

func help(update tgbotapi.Update) Answer {
	return Answer{Signature: "help"}
}
