package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secretkey")

func GenerateJWT(data []byte) (string, error) {
	claims := jwt.MapClaims{
		"id": string(data),
		"exp": time.Now().Add(time.Minute * 60).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	bearer := ctx.Request.Header.Get("Bearer")

	if bearer == "" {
		return nil, errors.New("bearer token required")
	}

	token, _ := jwt.Parse(bearer, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cannot parse jwt token")
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("can't get json data or token invalid")
	}

	return token.Claims.(jwt.MapClaims), nil
}