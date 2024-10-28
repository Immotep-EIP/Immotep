package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"immotep/backend/database"
	_ "immotep/backend/docs"
	"immotep/backend/router"
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

func mainFunc() int {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return 1
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		return 1
	}
	defer func() { _ = db.Client.Disconnect() }()

	err = router.Routes().Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(mainFunc())
}
