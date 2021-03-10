package main

var defaultLanguage string
var messages map[string]map[string]string

func defineMessages() {
	defaultLanguage = "en"

	messages = make(map[string]map[string]string)

	messages["en"] = make(map[string]string)
	messages["ru"] = make(map[string]string)

	messages["en"][":help"] = "Wanna some help? Go read documentation dude"
	messages["en"][":specify_trip"] = "You should specify trip name.\n\nExample:\n/new Vacation in Germany"
	messages["en"][":already_has_active_trip"] = "You already have an active trip in this chat"
	messages["en"][":no_active_trips"] = "You have no active trips in this chat yet"
	messages["en"][":you_are_already_in"] = "You're already in the trip!"
	messages["en"][":you_are_in"] = "Done, you're in!"
	messages["en"][":add_usage"] = "Usage: /add <sum> <description>"
	messages["en"][":you_are_not_a_trip_member"] = "You're not a trip member, try /join"
	messages["en"][":expense_created"] = "Expense created"
	messages["en"][":default_answer"] = "Oh, man, I don't know what are you talking about!"
	messages["en"][":use_group"] = "Add bot to the group chat to use it"

	messages["ru"][":use_group"] = "Чтобы использовать бота добавьте его в чат"
}
