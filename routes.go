package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Route struct {
	Prefix string
	Description string
	Callback func(update tgbotapi.Update) string
}

type RoutesRegistry struct {
	Routes []Route
}

func processUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	answer := rr.determineMethod(update.Message.Text)(update)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}


// TODO add descriptions
func (rr *RoutesRegistry) registerRoutes() {
	rr.addRoute("/help", "", help)
	rr.addRoute("/new", "", createNewTrip)
	rr.addRoute("/join", "", addMember)
	rr.addRoute("/members", "", getMembers)
	rr.addRoute("/add", "", addDebt)
}

func (rr *RoutesRegistry) addRoute(prefix string, description string, callback func(update tgbotapi.Update) string) {
	rr.Routes = append(rr.Routes, Route{prefix, description, callback})
}

func (rr *RoutesRegistry) determineMethod(message string) func(update tgbotapi.Update) string {
	for _, route := range rr.Routes  {
		if strings.HasPrefix(message, route.Prefix) {
			return route.Callback
		}
	}

	return defaultAnswer
}

func (rr *RoutesRegistry) setDescriptions() {
	// TODO
}