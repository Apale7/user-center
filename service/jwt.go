package service

import (
	"fmt"
	"os"
	"time"
	db_model "user_center/dal/db/model"
	"user_center/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	accessMaxAge  = 100
	refreshMaxAge = 60 * 60 * 24 * 7
)

var (
	accessSecret  string
	refreshSecret string
)

func init() {

	accessSecret, refreshSecret = getJWTConf()
	logrus.Info(accessSecret)
	if os.Getenv("ENV") == "dev" {
		accessSecret = "123456"
	}
}

func getJWTConf() (accessKey, refreshKey string) {
	viper.SetConfigName("jwt_conf")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper readInConfig error")
	}

	return viper.GetString("accessKey"), viper.GetString("refreshKey")
}

// createToken create token with uid and extra, expires after timeLength seconds.
func createToken(userID uint, userExtra db_model.UserExtra, timeLength int, key string) (tokenString string, exp int64, err error) {
	customClaims := model.CustomClaims{
		UserID: userID,
		Extra:  userExtra,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeLength) * time.Second).Unix(),
			Issuer:    "user_center",
		},
	}
	exp = customClaims.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err = token.SignedString([]byte(key))
	if err != nil {
		return
	}
	return
}

func CreateAccessToken(userID uint, userExtra db_model.UserExtra) (tokenString string, exp int64, err error) {
	return createToken(userID, userExtra, accessMaxAge, accessSecret)
}

func CreateRefreshToken(userID uint, userExtra db_model.UserExtra) (tokenString string, exp int64, err error) {
	return createToken(userID, userExtra, refreshMaxAge, refreshSecret)
}

func ParseToken(tokenString string, isRefresh bool) (model.CustomClaims, error) {
	var key string
	if isRefresh {
		key = refreshSecret
	} else {
		key = accessSecret
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return model.CustomClaims{}, errors.WithStack(err)
	}
	if claims, ok := token.Claims.(model.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Printf("%+v", claims)
		return model.CustomClaims{}, errors.WithStack(err)
	}
}
