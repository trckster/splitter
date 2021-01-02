package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Route struct {
	Prefix string
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

	var rr RoutesRegistry

	rr.registerRoutes()

	answer := rr.determineMethod(update.Message.Text)(update)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}


func (rr *RoutesRegistry) registerRoutes() {
	rr.addRoute("/help", help)
	rr.addRoute("/new", createNewTrip)
}

func (rr *RoutesRegistry) addRoute(prefix string, callback func(update tgbotapi.Update) string) {
	rr.Routes = append(rr.Routes, Route{prefix, callback})
}

func (rr *RoutesRegistry) determineMethod(message string) func(update tgbotapi.Update) string {
	for _, route := range rr.Routes  {
		if strings.HasPrefix(message, route.Prefix) {
			return route.Callback
		}
	}

	return defaultAnswer
}