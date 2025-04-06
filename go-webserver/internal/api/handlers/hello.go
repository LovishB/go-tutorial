package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
helloHandler is a simple HTTP handler that responds with "Hello World"
*/
func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
