package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"keyz/backend/docs"
	"keyz/backend/router"
	"keyz/backend/services"
)

//	@title			Keyz API
//	@version		1.0
//	@description	This is the API used by the Keyz application.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Mazettt
//	@contact.email	martin.d-herouville@epitech.eu

//	@host		localhost:3001
//	@BasePath	/v1

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer ` prefix, e.g. "Bearer abcde12345".

func checkEnvVar(key string) bool {
	value, check := os.LookupEnv(key)
	if !check || value == "" {
		log.Println("No " + key + " found in env")
		return false
	}
	return true
}

func mainFunc() int {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: failed loading .env file")
	}

	envVars := []string{
		"PORT",
		"PUBLIC_URL",
		"WEB_PUBLIC_URL",
		"SHOWCASE_PUBLIC_URL",
		"DATABASE_URL",
		"SECRET_KEY",
		"OPENAI_API_KEY",
		"BREVO_API_KEY",
	}
	for _, key := range envVars {
		if !checkEnvVar(key) {
			return 1
		}
	}

	docs.SwaggerInfo.Host = os.Getenv("PUBLIC_URL")

	db, err := services.ConnectDB()
	if err != nil {
		log.Println(err)
		return 1
	}
	log.Println("Connected to database, starting server...")
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
