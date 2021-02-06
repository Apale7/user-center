package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserExtraID uint
	Username    string `gorm:"uniqueIndex;size:64"`
	Password    string `gorm:"size:64"`
}

func (u *User) TableName() string {
	return "user"
}
