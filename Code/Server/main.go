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
		log.Println("WARNING: failed loading .env file")
	}

	port, check := os.LookupEnv("PORT")
	if !check {
		log.Println("No PORT found in env")
		return 1
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err)
		return 1
	}
	defer func() { _ = db.Client.Disconnect() }()

	log.Println("Connected to database, starting server...")

	err = router.Routes().Run(":" + port)
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(mainFunc())
}
