package jwt_auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Sp3ctroid/ED-API/types"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(key string, user types.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(user.ID),
		"sub":     user.Is_admin,
		"iss":     "API",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateToken(key string, tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token, nil

}
