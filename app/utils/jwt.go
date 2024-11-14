// utils/jwt.go
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a custom claims type
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Secret key used for signing tokens
var secretKey = []byte("your_secret_key") // Change this to a strong secret

// GenerateToken creates a new JWT token for a given user ID
func GenerateToken(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken checks the validity of the JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
