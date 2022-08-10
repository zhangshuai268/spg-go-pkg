package jwt

import (
	"errors"
	"github.com/go-pay/gopay/pkg/jwt"
)

//jwt,token验证
var jwtKey []byte

type MyClaims struct {
	jwt.StandardClaims
}

func GetToken(secret string, claims jwt.Claims) (string, error) {
	appSecret := secret
	jwtKey = []byte(appSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string, secret string, claims jwt.Claims) (error) {

	jwtKey = []byte(secret)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == nil && token != nil {
		if token.Valid {
			return nil
		} else {
			return errors.New("token解析失败")
		}
	}
	return err
}
