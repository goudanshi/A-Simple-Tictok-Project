package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

type JWT struct {
	privateKey []byte `json:"private_key"`
}

func NewJWT(privateKey []byte) *JWT {
	return &JWT{privateKey: privateKey}
}

func (j *JWT) CreateToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims) // 第一个参数是加密算法
	return token.SignedString(j.privateKey)
}

func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return j.privateKey, nil
	})
	if err != nil {
		//Logger.Error("parse token information err: " + err.Error())
		return nil, err
	}
	if token == nil {
		return nil, errors.New("parse token information error")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
