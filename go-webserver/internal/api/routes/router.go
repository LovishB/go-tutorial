package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	helloHandler "example.com/webserver/internal/api/handlers/hello"
	loginHandler "example.com/webserver/internal/api/handlers/login"
	walletHandler "example.com/webserver/internal/api/handlers/wallet"
	websocketHandler "example.com/webserver/internal/api/handlers/websocket"

	"example.com/webserver/internal/api/middleware"
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

	// Public routes
	router.GET("/hello", helloHandler.HelloHandler)
	router.POST("/login", loginHandler.Login)
	router.GET("/ws", websocketHandler.WsHandler)
	router.GET("/broadcast", websocketHandler.WsBroadcastHandler)

	// Protected routes
	router.GET("/wallet/:userId", middleware.AuthMiddleware(), walletHandler.GetWalletBalanceHandler)
	router.POST("/wallet", middleware.AuthMiddleware(), walletHandler.CreateWalletHandler)
	router.PUT("/wallet/:userId", middleware.AuthMiddleware(), walletHandler.UpdateWalletHandler)
}
