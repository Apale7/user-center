package handler

import (
	"context"
	"user_center/dal/db"
	"user_center/dal/db/model"
	"user_center/dto"
	user_center "user_center/proto/user-center"
)

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	err = db.Register(ctx, dto.RPCUser2ModelUser(req.User), model.UserExtra{})
	if err != nil {
		return
	}
	resp = &user_center.RegisterResponse{}
	return resp, nil 
}
