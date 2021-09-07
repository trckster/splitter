package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"regexp"
)

type Route struct {
	Prefix      string
	Description string
	Callback    func(update tgbotapi.Update) Answer
}

func processMessage(update tgbotapi.Update) {
	log.Printf("Message: [%s] %s", update.Message.From.UserName, update.Message.Text)

	var answer Answer

	if stateMethod, state := getCurrentStateMethod(update); stateMethod != nil {
		answer = stateMethod(update, state)
	} else {
		answer = getReplyMessage(update)
	}

	answer.send(update.Message)
}

func processCallback(update tgbotapi.Update) {
	log.Printf("Callback: [%s] %s", update.CallbackQuery.From.UserName, update.CallbackQuery.Data)

	var answer Answer

	if update.CallbackQuery.Data == "ask-name" {
		answer = askName(update)
	}

	bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))

	answer.send(update.CallbackQuery.Message)
}

func getReplyMessage(update tgbotapi.Update) Answer {
	if update.Message.Chat.ID >= 0 {
		return Answer{Signature: "use-group"}
	}

	command := getCommandByMessage(update.Message.Text)

	return rr.determineMethod(command)(update)
}

func getCommandByMessage(message string) string {
	re := regexp.MustCompile(`^/[a-zA-Z]*`)

	return string(re.Find([]byte(message)))
}
