package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Login string `json:"login"`
	Exp   int64  `json:"exp"`
}

func (t TokenClaims) Valid() error {
	if time.Now().Unix() > t.Exp {
		return errors.New("token has expired")
	}
	return nil
}

func New(tokenKey string, claims TokenClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tokenKey))
}

func ExtractClaims(tokenKey string, tokenStr string) (*TokenClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
