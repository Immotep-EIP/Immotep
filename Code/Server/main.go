package main

import (
	"immotep/backend/database"
	_ "immotep/backend/docs"
	"immotep/backend/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//	@title			Immotep API
//	@version		1.0
//	@description	This is the API used by the Immotep application.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Mazettt
//	@contact.email	martin.d-herouville@epitech.eu

//	@host		localhost:3001
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()
	defer db.Client.Disconnect()

	router.Routes().Run(":" + os.Getenv("PORT"))
}
