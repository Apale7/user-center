package service

import (
	"fmt"
	"time"
	db_model "user_center/dal/db/model"
	"user_center/model"

	"github.com/dgrijalva/jwt-go"
)

const (
	maxAge = 60 * 60 * 24 * 7
	
)

func CreateToken(userID uint, userExtra db_model.UserExtra, secret string) (tokenString string, err error) {
	customClaims := model.CustomClaims{
		UserID: userID,
		Extra:  userExtra,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
			Issuer:    "user_center",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, customClaims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}
	return
}

func ParseToken(tokenString string, secret string) (model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(model.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return model.CustomClaims{}, err
	}
}
