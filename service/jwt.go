package service

import (
	"fmt"
	"os"
	"time"
	db_model "user_center/dal/db/model"
	"user_center/model"

	"github.com/Apale7/common/constdef"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	accessMaxAge  = 60
	refreshMaxAge = 60 * 60 * 24 * 7
)

var (
	secret string
)

func init() {

	secret = getJWTConf()
	logrus.Info(secret)
	if os.Getenv("ENV") == "dev" {
		secret = "123456"
	}
}

func getJWTConf() string {
	viper.SetConfigName("jwt_conf")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper readInConfig error")
	}

	return viper.GetString(constdef.UserCenterSecretKey)
}

// createToken create token with uid and extra, expires after timeLength seconds.
func createToken(userID uint, userExtra db_model.UserExtra, timeLength int) (tokenString string, err error) {
	customClaims := model.CustomClaims{
		UserID: userID,
		Extra:  userExtra,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeLength) * time.Second).Unix(),
			Issuer:    "user_center",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}
	return
}

func CreateAccessToken(userID uint, userExtra db_model.UserExtra) (tokenString string, err error) {
	return createToken(userID, userExtra, accessMaxAge)
}

func CreateRefreshToken(userID uint, userExtra db_model.UserExtra) (tokenString string, err error) {
	return createToken(userID, userExtra, refreshMaxAge)
}

func ParseToken(tokenString string) (model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return model.CustomClaims{}, err
	}
	if claims, ok := token.Claims.(model.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Printf("%+v", claims)
		return model.CustomClaims{}, err
	}
}
