package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/nerd500/axios-cp-wing/handlers" 
)

func main() {
	router := gin.Default()

	
	router.GET("/", handlers.Ping)

	port := 8080
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	
}
