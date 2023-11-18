package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nerd500/axios-cp-wing/handlers"
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

	startServer(&apiHandler)

}
