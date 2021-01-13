package main

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	Name string
	OwnerID int
	ChatID int64
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

type Expense struct {
	gorm.Model
	Description string
	Amount float32
	TripID uint
	PayingID uint
	TripMember TripMember `gorm:"foreignKey:PayingID"`
	Trip Trip
}

type Debt struct {
	gorm.Model
	DebtorID int64
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