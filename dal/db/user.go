package db

import (
	"context"
	"user_center/dal/db/model"
	"errors"
	ex_errors"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Login(ctx context.Context, user model.User) (res model.User, err error) {
	err = db.WithContext(ctx).Order("username").Where(user).Take(&res).Error
	return
}

func Register(ctx context.Context, user model.User, extra model.UserExtra) (err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(&extra).Error; err != nil {
			err = ex_errors.Errorf("Create userExtra error: %v", err)
			return
		}
		logrus.Infof("Create UserExtra: %+v", extra)
		if extra.ID <= 0 {
			return errors.New("id is 0")
		}
		user.UserExtraID = extra.ID
		result := tx.FirstOrCreate(&user)
		if err = result.Error; err != nil {
			err = ex_errors.Errorf("Create user error: %v", err)
			return
		}
		if result.RowsAffected == 0 {//插入条数为0，用户名冲突
			err = ex_errors.New("用户名已存在")
			return
		}
		return
	})
	if err != nil {
		logrus.Warnf("Register error: %+v", err)
	}
	return err
}
