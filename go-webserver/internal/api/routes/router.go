package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	handler "example.com/webserver/internal/api/handlers"
)

/*
Router receives incoming HTTP requests
determines which handler should process request based on the URL path
@param router *gin.Engine
@return void
*/
func SetupRoutes(router *gin.Engine) {

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins - customize this in production
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Definning routes here
	router.GET("/hello", handler.HelloHandler)
}
