package db

import (
	"context"
	"user_center/dal/db/model"
)

func CreateGroup(ctx context.Context, groupInfo model.Group) (model.Group, error) {
	err := db.WithContext(ctx).Create(&groupInfo).Error
	if err == nil && groupInfo.ID <= 0 {
		panic("create group error")
	}
	if err != nil {
		return model.Group{}, err
	}
	return groupInfo, nil
}

func GetGroup(ctx context.Context, groupReqParams model.Group) (group []model.Group, err error) {
	db := db.WithContext(ctx).Model(&model.Group{}).Select("id, owner_id, group_name, created_at")
	if groupReqParams.ID > 0 {
		db = db.Where("id = ?", groupReqParams.ID)
	}
	if groupReqParams.OwnerID > 0 {
		db = db.Where("owner_id = ?", groupReqParams.OwnerID)
	}
	if len(groupReqParams.GroupName) > 0 {
		db = db.Where("group_name like ?", groupReqParams.GroupName+"%")
	}
	err = db.Find(&group).Error
	return
}

func GetGroupMembers(ctx context.Context, groupID uint) (users []model.UserExtra, err error) {
	var userIDs []uint
	db := db.WithContext(ctx)
	err = db.Model(&model.UserGroup{}).Where("group_id = ?", groupID).Pluck("user_id", &userIDs).Error
	if err != nil {
		return
	}
	err = db.Model(&model.UserExtra{}).Where("user_id in ?", userIDs).Find(&users).Error
	if err != nil {
		return
	}
	return
}

func JoinGroup(ctx context.Context, ug model.UserGroup) (err error) {
	db := db.WithContext(ctx)
	err = db.Model(&model.UserGroup{}).Create(&ug).Error
	return
}

func ExitGroup(ctx context.Context, ug model.UserGroup) (err error) {
	db := db.WithContext(ctx)
	err = db.Model(&model.UserGroup{}).Where(&ug).Delete(&ug).Error
	return err
}
