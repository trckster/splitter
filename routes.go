package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type Route struct {
	Prefix      string
	Description string
	Callback    func(update tgbotapi.Update) (string, interface{})
}

type RoutesRegistry struct {
	Routes []Route
}

func processUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	signature, markdown := getReplyMessage(update)

	language := determineLanguage(update)

	answer := messages[language][signature]

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = markdown

	bot.Send(msg)
}

func getReplyMessage(update tgbotapi.Update) (string, interface{}) {
	if update.Message.Chat.ID >= 0 {
		return ":use_group", nil
	}

	return rr.determineMethod(update.Message.Text)(update)
}

func determineLanguage(update tgbotapi.Update) string {
	trip, err := getCurrentTrip(update)

	if err == nil {
		return trip.Language
	}

	userLanguage := update.Message.From.LanguageCode

	if messages[userLanguage] != nil {
		return userLanguage
	}

	return defaultLanguage
}

func (rr *RoutesRegistry) registerRoutes() {
	rr.addRoute("/help", "Background information", help)
	rr.addRoute("/new", "Creates new trip", createNewTrip)
	rr.addRoute("/join", "Allows you to join current trip", addMember)
	rr.addRoute("/members", "See all trip members", getMembers)
	rr.addRoute("/add", "Add a debt", addDebt)
}

func (rr *RoutesRegistry) addRoute(prefix string, description string, callback func(update tgbotapi.Update) (string, interface{})) {
	rr.Routes = append(rr.Routes, Route{prefix, description, callback})
}

func (rr *RoutesRegistry) determineMethod(message string) func(update tgbotapi.Update) (string, interface{}) {
	for _, route := range rr.Routes {
		if strings.HasPrefix(message, route.Prefix) {
			return route.Callback
		}
	}

	return defaultAnswer
}

func (rr *RoutesRegistry) setDescriptions() {
	var commands []tgbotapi.BotCommand

	for _, route := range rr.Routes {
		commands = append(commands, tgbotapi.BotCommand{
			Command:     route.Prefix,
			Description: route.Description,
		})
	}

	err := bot.SetMyCommands(commands)

	if err != nil {
		panic(err)
	}
}
