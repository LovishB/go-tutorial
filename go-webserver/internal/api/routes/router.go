package routes

import (
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
	// Definning routes here
	router.GET("/hello", handler.HelloHandler)
}
