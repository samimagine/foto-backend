package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(userID uint, role string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   string(userID),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}