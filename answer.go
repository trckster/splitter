package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type Answer struct {
	Signature  string
	Message    string
	Language   string
	Keyboard   tgbotapi.InlineKeyboardMarkup
	Parameters map[string]string
}

func (answer *Answer) constructBotMessage(message *tgbotapi.Message) tgbotapi.MessageConfig {
	answer.determineLanguage(message)
	answer.prepareMessage()

	msg := tgbotapi.NewMessage(message.Chat.ID, answer.Message)
	msg.ReplyToMessageID = message.MessageID

	if len(answer.Keyboard.InlineKeyboard) != 0 {
		answer.prepareKeyboard()
		msg.ReplyMarkup = answer.Keyboard
	}

	return msg
}

func (answer *Answer) determineLanguage(message *tgbotapi.Message) {
	trip, err := getCurrentTrip(message)

	if err == nil {
		answer.Language = trip.Language
		return
	}

	userLanguage := message.From.LanguageCode

	if messages[userLanguage] != nil {
		answer.Language = userLanguage
		return
	}

	answer.Language = defaultLanguage
}

func (answer *Answer) prepareKeyboard() {
	var newKeyboardMarkup [][]tgbotapi.InlineKeyboardButton

	for _, row := range answer.Keyboard.InlineKeyboard {
		var newRow []tgbotapi.InlineKeyboardButton

		for _, element := range row {
			newButton := element
			newButton.Text = substituteBindings(keyboardsTexts[answer.Language][element.Text], answer.Parameters)
			newRow = append(newRow, newButton)
		}

		newKeyboardMarkup = append(newKeyboardMarkup, newRow)
	}

	answer.Keyboard.InlineKeyboard = newKeyboardMarkup
}

func (answer *Answer) prepareMessage() {
	answer.Message = messages[answer.Language][answer.Signature]

	answer.Message = substituteBindings(answer.Message, answer.Parameters)
}

func (answer *Answer) send(incomingMessage *tgbotapi.Message) {
	message := answer.constructBotMessage(incomingMessage)

	bot.Send(message)
}

func substituteBindings(string string, parameters map[string]string) string {
	for key, name := range parameters {
		string = strings.ReplaceAll(string, key, name)
	}

	return string
}
