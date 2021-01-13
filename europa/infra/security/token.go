package security

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// Claims defines JWT data structure
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken generate jwt with email as data
func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// ValidateToken validate jwt and return email
func ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	email := ""

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email = claims["email"].(string)
	} else {
		log.Printf("Error: %s", err)
	}

	return email, nil
}
