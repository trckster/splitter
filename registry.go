package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type RoutesRegistry struct {
	Routes map[string]Route
}

func (rr *RoutesRegistry) init() {
	rr.Routes = make(map[string]Route)
}

func (rr *RoutesRegistry) registerRoutes() {
	rr.addRoute("/start", "Start", start)
	rr.addRoute("/help", "Background information", help)
	rr.addRoute("/join", "Allows you to join current trip", addMember)
	rr.addRoute("/members", "See all trip members", getMembers)
	rr.addRoute("/add", "Add a debt", addDebt)
}

func (rr *RoutesRegistry) addRoute(prefix string, description string, callback func(update tgbotapi.Update) Answer) {
	rr.Routes[prefix] = Route{prefix, description, callback}
}

func (rr *RoutesRegistry) determineMethod(message string) func(update tgbotapi.Update) Answer {
	if route, ok := rr.Routes[message]; ok {
		return route.Callback
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
