package dto

import (
	"user_center/dal/db/model"
	user_center "user_center/proto/user-center"

	"gorm.io/gorm"
)

func ToGRPCGroup(g *model.Group) *user_center.Group {
	if g == nil {
		return &user_center.Group{}
	}
	return &user_center.Group{
		Id:        uint32(g.ID),
		CreatedAt: g.CreatedAt.Local().Unix(),
		UpdatedAt: g.UpdatedAt.Local().Unix(),
		OwnerId:   g.OwnerID,
		GroupName: g.GroupName,
	}
}

func ToModelGroup(g *user_center.Group) *model.Group { //写入数据时才会用ToModelGroup, time字段可以不转
	if g == nil {
		return &model.Group{}
	}
	return &model.Group{
		Model: gorm.Model{
			ID: uint(g.Id),
		},
		OwnerID:   g.OwnerId,
		GroupName: g.GroupName,
	}
}
