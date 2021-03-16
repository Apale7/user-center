package handler

import (
	"context"
	"fmt"
	"testing"
	user_center "user_center/proto/user-center"
	"user_center/service"

	"github.com/sirupsen/logrus"
)

var userCtx = context.Background()

func TestLogin(t *testing.T) {
	req := &user_center.LoginRequest{
		Username: "apale",
		Password: "123465",
	}
	resp, err := Login(userCtx, req)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp)

	claims, err := service.ParseToken(resp.AccessToken, false)
	if err != nil {
		logrus.Errorln(err)
		t.FailNow()
	}
	fmt.Printf("%+v", claims)
}

func TestCheckToken(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user_center.CheckTokenRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.CheckTokenResponse
		wantErr  bool
	}{
		{
			name: "check token",
			args: args{ctx: ctx, req: &user_center.CheckTokenRequest{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIkV4dHJhIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTAyLTI3VDIxOjI5OjA5KzA4OjAwIiwiVXBkYXRlZEF0IjoiMjAyMS0wMi0yN1QyMToyOTowOSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiVXNlcklEIjoxLCJOaWNrbmFtZSI6IkFwYWxlIiwiUGhvbmVOdW1iZXIiOiIxMjM0NTY3ODkwIiwiRW1haWwiOiIiLCJBdmF0YXJVUkwiOiIifSwiZXhwIjoxNjE1OTE4NjE5LCJpc3MiOiJ1c2VyX2NlbnRlciJ9.U0KHsfJBUygMXTpTRGV4BnCaoieIYnn9oe8Uz3UzM28",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckToken(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("CheckToken() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}
