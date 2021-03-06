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

// haveMe 为false时返回不包括memberID的group，用于加入团队
func GetGroup(ctx context.Context, groupInfo model.Group, memberID uint32, haveMe bool) (group []model.Group, err error) {
	mainDB := db.WithContext(ctx).Model(&model.Group{}).Select("id, owner_id, group_name, created_at")
	if groupInfo.ID > 0 {
		mainDB = mainDB.Where("id = ?", groupInfo.ID)
	}
	if groupInfo.OwnerID > 0 {
		mainDB = mainDB.Where("owner_id = ?", groupInfo.OwnerID)
	}
	if memberID > 0 { //传了memberID则按member所在的group进行查询
		var groupIDs []uint32
		subDB := db.WithContext(ctx).Model(&model.UserGroup{})
		subDB.Select("group_id").Where("user_id=?", memberID).Pluck("group_id", &groupIDs) //获取member所在的所有group

		if haveMe {
			mainDB = mainDB.Where("id in ?", groupIDs)
		} else if len(groupIDs) > 0 {
			mainDB = mainDB.Where("id not in ?", groupIDs)
		}
	}

	if len(groupInfo.GroupName) > 0 {
		mainDB = mainDB.Where("group_name like ?", groupInfo.GroupName+"%")
	}
	err = mainDB.Find(&group).Error
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
	err = db.Unscoped().Model(&model.UserGroup{}).Where(&ug).Delete(&ug).Error
	return err
}
