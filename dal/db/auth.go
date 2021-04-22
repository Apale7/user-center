package db

import (
	"context"
	"user_center/dal/db/model"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func GetAuthList(ctx context.Context, userID uint) (authList []string, err error) {
	roleIDs, err := getRoles(ctx, userID)
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}

	resourceIDs, err := getResourceIDs(ctx, roleIDs)
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}

	authList, err = getResourceNames(ctx, resourceIDs)
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}

	return
}

func getRoles(ctx context.Context, userID uint) (roleIDs []uint, err error) {
	db := db.WithContext(ctx)
	err = db.Model(&model.UserRole{}).Select("role_id").Where("user_id = ?", userID).Pluck("role_id", &roleIDs).Error
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}
	return roleIDs, nil
}

func getResourceIDs(ctx context.Context, roleIDs []uint) (resourceIDs []uint, err error) {
	db := db.WithContext(ctx)
	err = db.Model(&model.RoleResource{}).Select("resource_id").Where("role_id in ?", roleIDs).Pluck("resource_id", &resourceIDs).Error

	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}
	return resourceIDs, nil
}

func getResourceNames(ctx context.Context, resourceIDs []uint) (resourceNames []string, err error) {
	db := db.WithContext(ctx)
	err = db.Model(&model.Resource{}).Select("name").Where("id in ?", resourceIDs).Pluck("name", &resourceNames).Error
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}
	return resourceNames, nil
}
