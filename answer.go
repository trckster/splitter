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

//type KeyboardElement struct {
//	Signature  string
//	Data       string
//	Parameters []string
//}

func (answer *Answer) constructBotMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	answer.determineLanguage(update)


	answer.prepareMessage()

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.Message)
	msg.ReplyToMessageID = update.Message.MessageID

	if len(answer.Keyboard.InlineKeyboard) != 0 {
		answer.prepareKeyboard()
		msg.ReplyMarkup = answer.Keyboard
	}

	return msg
}

func (answer *Answer) determineLanguage(update tgbotapi.Update) {
	trip, err := getCurrentTrip(update)

	if err == nil {
		answer.Language = trip.Language
		return
	}

	userLanguage := update.Message.From.LanguageCode

	if messages[userLanguage] != nil {
		answer.Language = userLanguage
		return
	}

	answer.Language = defaultLanguage
}

func (answer *Answer) prepareKeyboard() {
	for i, row := range answer.Keyboard.InlineKeyboard {
		for j, element := range row {
			element.Text = keyboardsTexts[answer.Language][element.Text]
			answer.Keyboard.InlineKeyboard[i][j].Text = substituteBindings(element.Text, answer.Parameters)
		}
	}
}

func (answer *Answer) prepareMessage() {
	answer.Message = messages[answer.Language][answer.Signature]

	answer.Message = substituteBindings(answer.Message, answer.Parameters)
}

func substituteBindings(string string, parameters map[string]string) string {
	for key, name := range parameters {
		string = strings.ReplaceAll(string, key, name)
	}

	return string
}
