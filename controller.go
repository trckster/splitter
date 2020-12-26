package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func help(update tgbotapi.Update) string {
	return "Wanna some help? Go read documentation dude"
}

func defaultAnswer(update tgbotapi.Update) string {
	return "Oh, man, I don't know what are you talking about!"
}