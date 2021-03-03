package handler

import (
	"context"
	"user_center/dal/db"
	"user_center/dal/db/model"
	"user_center/dto"
	"user_center/proto/base"
	user_center "user_center/proto/user-center"
	"user_center/service"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	user, err := db.GetUser(ctx, model.User{Username: req.Username})
	if err != nil {
		return
	}

	if user.Password != req.Password {
		return nil, errors.New("用户名或密码错误")
	}
	extra, err := db.GetUserExtra(ctx, user)
	if err != nil {
		return
	}

	accessToken, err := service.CreateAccessToken(user.ID, extra)
	if err != nil {
		return nil, errors.Errorf("create accessToken error, err: %+v", err)
	}
	refreshToken, err := service.CreateRefreshToken(user.ID, extra)
	if err != nil {
		return nil, errors.Errorf("create refreshToken error, err: %+v", err)
	}
	resp = &user_center.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return
}

func CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (resp *user_center.CheckTokenResponse, err error) {
	claims, err := service.ParseToken(req.Token)
	if err != nil {
		return
	}
	logrus.Infof("%s access", claims.Extra.Nickname)
	resp = &user_center.CheckTokenResponse{
		BaseResp: &base.BaseResp{StatusCode: 0},
	}
	return
}

func Refresh(ctx context.Context, req *user_center.RefreshRequest) (resp *user_center.RefreshResponse, err error) {
	claims, err := service.ParseToken(req.RefreshToken)
	if err != nil {
		return
	}
	logrus.Infof("%s refresh", claims.Extra.Nickname)
	accessToken, err := service.CreateAccessToken(claims.UserID, claims.Extra)
	if err != nil {
		return nil, errors.Errorf("create accessToken error, err: %+v", err)
	}
	refreshToken, err := service.CreateRefreshToken(claims.UserID, claims.Extra)
	if err != nil {
		return nil, errors.Errorf("create refreshToken error, err: %+v", err)
	}
	resp = &user_center.RefreshResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return
}
