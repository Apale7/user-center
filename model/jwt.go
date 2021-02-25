package model

import "github.com/dgrijalva/jwt-go"
import db_model "user_center/dal/db/model"

type CustomClaims struct {
	UserID uint
	Extra  db_model.UserExtra
	jwt.StandardClaims
}
