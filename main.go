package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nerd500/axios-cp-wing/client/email_client"
	"github.com/nerd500/axios-cp-wing/handlers/handlers"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

func setEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initialiseDB() database.Querier {
	db, err := database.InitialiseDatabase()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func startServer(apiHandler *handlers.Api) {
	router := handlers.SetupRoutes(apiHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func main() {

	setEnv()

	apiHandler := handlers.Api{}
	apiHandler.DB = initialiseDB()
	defer database.CloseDataBase()

	apiHandler.EmailClient = &email_client.EmailClientImpl{
		AppEmail: os.Getenv("APP_EMAIL"),
		Password: os.Getenv("APP_EMAIL_PASSWORD"),
	}
	startServer(&apiHandler)

}
