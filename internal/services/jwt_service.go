package services

import (
	"beyond/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID.String(),
		"name": user.Nome,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	})
	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {

	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		// ESSA LINHA É A CHAVE: Ela vai imprimir no terminal o motivo real
		fmt.Println("MOTIVO DO TOKEN INVÁLIDO:", err)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}
