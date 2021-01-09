package main

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	Name string
	OwnerId int
	ChatId int64
	Members []TripMember
}

type TripMember struct {
	gorm.Model
	UserID int
	Username string
	FirstName string
	TripID uint
	Trip Trip
}

type Debt struct {
	gorm.Model
	Description string
	DebtorId int64
	LenderId int64
	Amount float32
	IsClosed bool
}

func (trip *Trip) addMember(userID int, username string, name string) *TripMember {
	member := TripMember {
		UserID: userID,
		TripID: trip.ID,
		Username: username,
		FirstName: name,
	}

	db.Create(&member)

	return &member
}