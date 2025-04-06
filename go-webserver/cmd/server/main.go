package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"example.com/webserver/internal/api/routes"
)

func main() {
	fmt.Printf("Starting the Web Server...\n")
	// Initialize the router & setting up routes
	router := gin.Default()
	routes.SetupRoutes(gin.Default())

	// Start the server
	router.Run(":8080")
	log.Println("Server starting on :8080")
}
