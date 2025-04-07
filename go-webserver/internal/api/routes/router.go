package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	helloHandler "example.com/webserver/internal/api/handlers/hello"
	walletHandler "example.com/webserver/internal/api/handlers/wallet"
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
	router.GET("/hello", helloHandler.HelloHandler)

	router.GET("/wallet/:userId", walletHandler.GetWalletBalanceHandler)
	router.POST("/wallet", walletHandler.CreateWalletHandler)
	router.PUT("/wallet/:userId", walletHandler.UpdateWalletHandler)
}
