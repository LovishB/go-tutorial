package handler

import (
	"net/http"

	"example.com/webserver/internal/auth"
	"example.com/webserver/internal/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginRequest := model.LoginRequest{}
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "invalid_request",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// Simulate a successful login
	if loginRequest.Password != "hellogolang" {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "invalid_credentials",
			Message: "Invalid username or password",
		}
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(loginRequest.Username)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "token_generation_failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	response := model.LoginResponse{
		Username: loginRequest.Username,
		Token:    token,
	}
	c.JSON(http.StatusOK, response)
}
