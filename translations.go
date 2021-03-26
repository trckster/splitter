package main

var defaultLanguage string
var messages map[string]map[string]string
var keyboardsTexts map[string]map[string]string

func defineTranslations() {
	defaultLanguage = "en"

	defineMessages()
	defineKeyboardsTexts()
}

func defineMessages() {
	messages = make(map[string]map[string]string)

	messages["en"] = make(map[string]string)
	messages["ru"] = make(map[string]string)

	messages["en"]["help"] = "Wanna some help? Go read documentation dude"
	messages["en"]["already-has-active-trip"] = "You already have an active trip in this chat"
	messages["en"]["create-first-trip"] = "Create your first trip, so you can split expenses."
	messages["en"]["no-active-trips"] = "You have no active trips in this chat yet"
	messages["en"]["you-are-already-in"] = "You're already in the trip!"
	messages["en"]["you-are-in"] = "Done, you're in!"
	messages["en"]["add-usage"] = "Usage: /add <sum> <description>"
	messages["en"]["you-are-not-a-trip-member"] = "You're not a trip member, try /join"
	messages["en"]["expense-created"] = "Expense created"
	messages["en"]["default-answer"] = "Oh, man, I don't know what are you talking about!"
	messages["en"]["use-group"] = "Add bot to the group chat to use it"
	messages["en"]["new-trip"] = "Successfully created new trip: :trip_name"
	messages["en"]["hack"] = ":hack"

	messages["ru"]["use-group"] = "Чтобы использовать бота добавьте его в чат"

	// New messages

	messages["en"]["input-trip-name"] = "Enter new trip name, please:"
}

func defineKeyboardsTexts() {
	keyboardsTexts = make(map[string]map[string]string)

	keyboardsTexts["en"] = make(map[string]string)
	keyboardsTexts["ru"] = make(map[string]string)

	keyboardsTexts["en"]["create-trip"] = "Create new trip"

	keyboardsTexts["ru"]["create-trip"] = "Создать поездку"
}