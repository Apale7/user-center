package handler

import (
	"context"
	"fmt"
	"testing"
	user_center "user_center/proto/user-center"
	"user_center/service"

	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func TestLogin(t *testing.T) {
	req := &user_center.LoginRequest{
		Username: "apale",
		Password: "123465",
	}
	resp, err := Login(ctx, req)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp.Token)

	claims, err := service.ParseToken(resp.Token)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	fmt.Printf("%+v", claims)
}
