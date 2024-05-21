package helper

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwt_token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	jwt_token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, valid := t.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("unauthorized %v", t.Header["alg"])
		}

		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims, valid := jwt_token.Claims.(jwt.MapClaims)
	if !jwt_token.Valid && !valid {
		return nil, errors.New("token not valid")
	}

	return claims, nil
}
