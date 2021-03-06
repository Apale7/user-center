package service

import (
	"context"
	"fmt"
	"testing"
	"user_center/dal/db/model"
)

var ctx = context.Background()

func TestParseService(t *testing.T) {
	token, _, _ := CreateAccessToken(123456, model.UserExtra{})
	fmt.Println(token)

	if res, err := ParseToken(token, true); err != nil {
		t.FailNow()
	} else {
		fmt.Printf("%+v", res)
	}
}
