package main

type Trip struct {
	ID int `gorm:"primaryKey"`
	Name string
	OwnerId int
	ChatId int64
}

type Debt struct {
	ID int `gorm:"primaryKey"`
	Description string
	DebtorId int64
	LenderId int64
	Amount float32
	IsClosed bool
}
