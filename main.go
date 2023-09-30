package main

import (
	"log"
	"net/http"

	"github.com/nerd500/axios-cp-wing/routes"
)

func main() {

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
