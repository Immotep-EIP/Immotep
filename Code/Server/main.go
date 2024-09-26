package main

import (
	"immotep/backend/database"
	"immotep/backend/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()
	defer db.Client.Disconnect()

	router.Routes().Run(":" + os.Getenv("PORT"))
}
