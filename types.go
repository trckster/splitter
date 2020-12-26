package main

type Trip struct {
	Id int
	Name string
	OwnerId int64
	ChatId int64
}

type Debt struct {
	Id int
	Description string
	DebtorId int64
	LenderId int64
	Amount float32
	IsClosed bool
}

