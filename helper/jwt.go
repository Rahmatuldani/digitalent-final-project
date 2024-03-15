package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretkey")

func GenerateJWT(data []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(data),
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}