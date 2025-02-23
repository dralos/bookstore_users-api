package main

import (
	"log"

	"github.com/dralos/bookstore_users-api/app"
	"github.com/dralos/bookstore_users-api/datasources/mysql/users_db"
	"github.com/joho/godotenv"
)

func main() {
	// Code
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	if err := users_db.SetupDatabase(); err != nil {
		log.Fatalf("database setup error : %v", err)
	}
	app.StartApp()
}
