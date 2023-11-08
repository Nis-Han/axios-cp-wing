package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/routes"
)

func setEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initialiseDB() {
	if err := database.InitialiseDatabase(); err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	router := routes.SetupRoutes()

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

	initialiseDB()
	defer database.CloseDataBase()

	startServer()

}
