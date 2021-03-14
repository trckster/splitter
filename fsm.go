package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

type FSM struct {
	gorm.Model
	State  string
	UserID int
	ChatID int64
	Data   string
}

var AvailableFSMs = map[string][]string{
	"CreateTrip": {
		"save-name",
		"save-something-other",
	},
}

var StateToMethod = map[string]func(update tgbotapi.Update, state *FSM) Answer{
	"save-name":            testFirstState,
	"save-something-other": testSecondState,
}

func testFirstState(update tgbotapi.Update, state *FSM) Answer {
	state.next()

	return Answer{Signature: "hack", Parameters: map[string]string{":hack": "Your name saved. Something other?"}}
}

func testSecondState(update tgbotapi.Update, state *FSM) Answer {
	state.next()

	return Answer{Signature: "hack", Parameters: map[string]string{":hack": "Thank you. We saved all."}}
}

func initState(update tgbotapi.Update, FSMName string) {
	state := FSM{
		UserID: update.Message.From.ID,
		ChatID: update.Message.Chat.ID,
		State:  AvailableFSMs[FSMName][0],
	}

	db.Save(&state)
}

func (state *FSM) next() {
	nextState := getNextState(state.State)

	if nextState == "" {
		db.Delete(&state)
	} else {
		state.State = nextState

		db.Save(&state)
	}
}

func getNextState(currentState string) string {
	for _, states := range AvailableFSMs {
		for i, state := range states {
			if currentState == state {
				if i+1 == len(states) {
					return ""
				}

				return states[i+1]
			}
		}
	}

	return ""
}

func getCurrentStateMethod(update tgbotapi.Update) (func(update tgbotapi.Update, state *FSM) Answer, *FSM) {
	state, err := getCurrentState(update)

	if err != nil {
		return nil, nil
	}

	return StateToMethod[state.State], state
}

func getCurrentState(update tgbotapi.Update) (*FSM, error) {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	var state FSM

	err := db.Where("user_id", userID).Where("chat_id", chatID).First(&state).Error

	return &state, err
}
