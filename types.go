package main

import (
	"gorm.io/gorm"
)

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
	Amount int
	TripID uint
	PayingID uint
	TripMember TripMember `gorm:"foreignKey:PayingID"`
	Trip Trip
}

type Debt struct {
	gorm.Model
	ExpenseID uint
	DebtorID uint
	Amount uint
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

func (trip *Trip) addExpense(payingID int, amount int, description string) (*Expense, error) {
	var member TripMember

	result := db.Where("user_id", payingID).Where("trip_id", trip.ID).Find(&member)

	if result.Error != nil {
		return nil, newError("You're not a trip member, try /join")
	}

	expense := &Expense {
		Amount: amount,
		TripID: trip.ID,
		PayingID: member.ID,
		Description: description,
	}

	db.Create(expense)

	return expense, nil
}

func (expense *Expense) split(payingID int) {
	var members []TripMember

	db.Where("trip_id", expense.TripID).Where("user_id <> ?", payingID).Find(&members)

	debtSum := expense.Amount / (len(members) + 1)

	for _, member := range members {
		member.addDebt(expense.ID, uint(debtSum))
	}
}

func (member *TripMember) addDebt(expenseID uint, memberDebt uint) *Debt {
	debt := &Debt {
		ExpenseID: expenseID,
		DebtorID: member.ID,
		Amount: memberDebt,
		IsClosed: false,
	}

	db.Create(debt)

	return debt
}