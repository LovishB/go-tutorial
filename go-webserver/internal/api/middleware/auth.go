package middleware

import (
	"strings"

	"example.com/webserver/internal/auth"
	"example.com/webserver/internal/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc { //returning a function that takes a gin.Context and returns a gin.HandlerFunc
	return func(c *gin.Context) {
		//Get the authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errResponse := model.ErrorResponse{
				Message: "Authorization header is missing",
				Code:    "header_missing",
				Status:  "error",
			}
			c.JSON(401, errResponse)
			c.Abort() //stop the request
			return
		}

		// Extract the token from the header (Bearer format)
		parts := strings.Split(authHeader, " ")      //split the header by space, header format -> (Bearer <token>)
		if len(parts) != 2 || parts[0] != "Bearer" { //authHeader should have 2 parts and the first part should be "Bearer"
			errResponse := model.ErrorResponse{
				Message: "Invalid authorization header format",
				Code:    "invalid_header_format",
				Status:  "error",
			}
			c.JSON(401, errResponse)
			c.Abort() //stop the request
			return
		}

		//Validate the token
		tokenString := parts[1] //token is the second part of the header
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			errResponse := model.ErrorResponse{
				Message: err.Error(),
				Code:    "invalid_token",
				Status:  "error",
			}
			c.JSON(401, errResponse)
			c.Abort() //stop the request
			return
		}

		c.Set("user_id", claims.UserID) //set the user id in the context
		c.Next()                        //call the next handler
	}
}
