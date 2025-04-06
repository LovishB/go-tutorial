package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"example.com/webserver/internal/api/routes"
	"example.com/webserver/internal/config"
)

func main() {
	fmt.Printf("Starting the Web Server...\n")
	// Initialize the router & setting up routes
	router := gin.Default()
	routes.SetupRoutes(router)

	// Load the configuration
	config := config.LoadConfig()

	router.Run(":" + config.Port)
	log.Println("Server starting on :" + config.Port)
}
