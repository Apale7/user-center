package model

import "gorm.io/gorm"

type UserType int8

const (
	UserType_Student UserType = 1
	UserType_Teacher UserType = 2
	UserType_Admin   UserType = 3
)

type UserExtra struct {
	gorm.Model
	UserID      uint `gorm:"uniqueIndex"`
	UserType    UserType
	Nickname    string `gorm:"size:64"`
	PhoneNumber string `gorm:"size:15"`
	Email       string `gorm:"size:64"`
	AvatarURL   string `gorm:"size:255"`
}

func (u UserExtra) TableName() string {
	return "user_extra"
}
