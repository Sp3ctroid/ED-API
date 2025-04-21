package jwt_auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(key string, userID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(userID),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
