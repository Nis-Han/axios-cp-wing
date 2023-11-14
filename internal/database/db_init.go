package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DBInstance Querier

var conn *sql.DB

func InitialiseDatabase() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	conn, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	if err = conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	fmt.Printf("connected to database \n")

	DBInstance = New(conn)

	return nil
}

func CloseDataBase() {
	conn.Close()
}
