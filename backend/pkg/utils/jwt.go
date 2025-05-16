package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "gdzie_masz_plik_dotenv"
	}
	return []byte(key)
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
	return token.SignedString(getJWTKey())
}
func ParseTokenJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return getJWTKey(), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
