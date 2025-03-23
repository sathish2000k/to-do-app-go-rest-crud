package utils

import (
	"time"

	models "to-do-app/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("jwt_secret_key")

type Claims struct {
	UserId int `json:"user_id"`
	UserPassword string `json:"password"`
	UserRole models.Role `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userId int, userPassword string, userRole models.Role, expirationTime time.Time) (string, error) {
	claims := &Claims{
		UserId: userId,
		UserPassword: userPassword,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error ) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}