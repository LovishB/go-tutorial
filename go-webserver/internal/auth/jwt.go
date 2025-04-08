package auth

import (
	"errors"
	"time"

	"example.com/webserver/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

// jwt secret key
var jwtKey = []byte(config.LoadConfig().AuthKey)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 24 hrs
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}
	// Signing claim with secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{} //Initialise a claim object

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil { //error while validating
		return nil, err
	}

	if !token.Valid { //invalid token
		return nil, errors.New("invalid token")
	}

	return claims, nil //Validated
}
