package main

import (
	"fmt"
	"immotep/backend/database"
	"immotep/backend/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	port := app.Config.Port
	fmt.Printf("Serving app on port %s\n", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return srv.ListenAndServe()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.Client.Disconnect()
}

func main() {
	LoadEnv()
	InitDB()

	config := Config{
		Port: os.Getenv("PORT"),
	}
	app := &Application{
		Config: config,
	}

	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
