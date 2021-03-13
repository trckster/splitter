package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

type Route struct {
	Prefix      string
	Description string
	Callback    func(update tgbotapi.Update) Answer
}

type RoutesRegistry struct {
	Routes []Route
}

func processUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	answer := getReplyMessage(update)

	message := answer.constructBotMessage(update)

	bot.Send(message)
}

func getReplyMessage(update tgbotapi.Update) Answer {
	if update.Message.Chat.ID >= 0 {
		return Answer{Signature: "use-group"}
	}

	return rr.determineMethod(update.Message.Text)(update)
}

func (rr *RoutesRegistry) registerRoutes() {
	rr.addRoute("/start", "Start", start)
	rr.addRoute("/help", "Background information", help)
	rr.addRoute("/new", "Creates new trip", createNewTrip)
	rr.addRoute("/join", "Allows you to join current trip", addMember)
	rr.addRoute("/members", "See all trip members", getMembers)
	rr.addRoute("/add", "Add a debt", addDebt)
}

func (rr *RoutesRegistry) addRoute(prefix string, description string, callback func(update tgbotapi.Update) Answer) {
	rr.Routes = append(rr.Routes, Route{prefix, description, callback})
}

func (rr *RoutesRegistry) determineMethod(message string) func(update tgbotapi.Update) Answer {
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
