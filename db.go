package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func constructDataSourceName() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	login := os.Getenv("DB_LOGIN")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	result := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host,
		login,
		password,
		database,
		port)

	return result
}

func connectToDatabase() {
	dsn := constructDataSourceName()

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Connected.")
}

func migrateAllModels() {
	db.AutoMigrate(&Trip{}, &Debt{}, &TripMember{}, &Expense{}, &FSM{})
}
