package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	conn, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	
	fmt.Printf("connected to database \n")
	defer conn.Close()

	db := database.New(conn)
	router := routes.SetupRoutes(db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
