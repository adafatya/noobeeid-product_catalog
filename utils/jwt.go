package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// generate jwt token with 1d expiration
// param issuer string
// return jwt token string, error
func GenerateJWTToken(issuer string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	secret := []byte(secretKey)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Issuer:    issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		log.Println("error saat pembuatan token jwt: ", err)
		return "", err
	}

	return ss, nil
}
