package main

import (
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var db *gorm.DB

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	connectToDatabase()

	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		processUpdate(update)
	}
}