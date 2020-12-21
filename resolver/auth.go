package resolver

import (
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
