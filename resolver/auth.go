package resolver

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const exp = "24h"
const secret = "andrefajar"

func createToken(user string) (string, error) {
	exp, err := time.ParseDuration(exp)
	if err != nil {

		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(exp).Unix(),
	})
	s, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// when authorization is nil
	if tokenString == "" {
		return nil, fmt.Errorf("Missing Authorization Header")
	}

	// decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("There was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// when token is invalid
	if !token.Valid {
		return nil, fmt.Errorf("Token invalid.")
	}

	return token.Claims.(jwt.MapClaims), nil
}
