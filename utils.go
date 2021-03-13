package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type SplitterError struct {
	message string
}

func (e SplitterError) Error() string {
	return e.message
}

func newError(message string) *SplitterError {
	return &SplitterError{message: message}
}

func getCurrentTrip(update tgbotapi.Update) (Trip, error) {
	var trip Trip

	record := db.Where("chat_id", update.Message.Chat.ID).First(&trip)

	if record.Error != nil {
		return trip, newError("no-active-trips")
	}

	return trip, nil
}

func hasCurrentTrip(update tgbotapi.Update) bool {
	_, err := getCurrentTrip(update)

	return err == nil
}
