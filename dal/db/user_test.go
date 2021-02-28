package db

import (
	"context"
	"fmt"
	"testing"
	"user_center/dal/db/model"

	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func TestGetUser(t *testing.T) {
	user, err := GetUser(ctx, model.User{Username: "apale"})
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	logrus.Infof("%+v", user)
}

func TestRegister(t *testing.T) {
	user := model.User{Username: "apale", Password: "123465"}
	extra := model.UserExtra{
		Nickname:    "Apale",
		PhoneNumber: "1234567890",
	}
	err := CreateUserWithExtra(ctx, user, extra)
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
}

func TestGetExtra(t *testing.T) {
	user := model.User{}
	user.ID = 1
	extra, err := GetUserExtra(ctx, user)
	if err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	fmt.Printf("%+v", extra)
}
