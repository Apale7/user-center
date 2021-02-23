package handler

import (
	"context"
	"errors"
	"user_center/dal/db"
	"user_center/dal/db/model"
	"user_center/dto"
	user_center "user_center/proto/user-center"
)

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	err = db.CreateUserWithExtra(ctx, dto.RPCUser2ModelUser(req.User), model.UserExtra{})
	if err != nil {
		return
	}
	resp = &user_center.RegisterResponse{}
	return resp, nil
}

func Login(ctx context.Context, req *user_center.LoginRequest) (resp *user_center.LoginResponse, err error) {
	user, err := db.GetUser(ctx, req.Username)
	if err != nil {
		return
	}

	if user.Password != req.Password {
		return nil, errors.New("用户名或密码错误")
	}
	resp = &user_center.LoginResponse{}
	return
}

func getToken(user model.User) {
	
}