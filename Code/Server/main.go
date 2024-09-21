package main

import (
	"immotep/backend/database"
	"immotep/backend/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *database.PrismaDB {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	return db
}

func main() {
	LoadEnv()

	db := ConnectDB()
	defer db.Client.Disconnect()

	router.Routes().Run(":" + os.Getenv("PORT"))
}
