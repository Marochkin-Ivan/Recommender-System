package storage

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

func GenerateToken(id string) string {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(30 * time.Hour * 24).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("secret")) // salt for jwt
	if err != nil {
		log.Println(err)
		return ""
	}

	return token
}

func CheckToken(token string) (Claims, error) {
	cl := Claims{}
	t, err := jwt.ParseWithClaims(token, &cl, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return "", errors.New("bad method")
		}
		return []byte("secret"), nil
	})
	if err != nil || !t.Valid {
		return Claims{}, err
	}

	return cl, nil
}
