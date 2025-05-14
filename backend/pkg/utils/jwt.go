package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret_key") //TODO: .env + parser

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ParseTokenJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
