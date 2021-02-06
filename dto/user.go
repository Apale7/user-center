package dto

import (
	"user_center/dal/db/model"
	user_center "user_center/proto/user-center"

	"gorm.io/gorm"
)

func RPCUser2ModelUser(user *user_center.User) model.User {
	return model.User{
		Username: user.Username,
		Password: user.Password,
		Model: gorm.Model{
			ID: uint(user.Id),
		},
	}
}

func ModelUser2RPCUser(user *model.User) user_center.User {
	return user_center.User{
		Id:       int32(user.ID),
		Username: user.Username,
		Password: user.Password,
	}
}
