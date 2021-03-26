package main

import (
	"encoding/json"
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
		"receive-name-and-create-trip",
	},
}

var StateToMethod = map[string]func(update tgbotapi.Update, state *FSM) Answer{
	"receive-name-and-create-trip": saveName,
}

func initState(from *tgbotapi.User, chat *tgbotapi.Chat, FSMName string) {
	state := FSM{
		UserID: from.ID,
		ChatID: chat.ID,
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

func (state *FSM) getData() map[string]interface{} {
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(state.Data), &data); err != nil {
		panic(err)
	}

	return data
}

func (state *FSM) addData(key string, additionalData interface{}) *FSM {
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(state.Data), &data); err != nil {
		data = make(map[string]interface{})
	}

	data[key] = additionalData

	dataAsBytes, _ := json.Marshal(data)

	state.Data = string(dataAsBytes)

	return state
}

func (state *FSM) saveData() {
	db.Save(state)
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
