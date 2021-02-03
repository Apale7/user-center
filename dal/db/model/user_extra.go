package model

import "gorm.io/gorm"

type UserExtra struct {
	gorm.Model
	Nickname    string `gorm:"size:64"`
	PhoneNumber string `gorm:"size:15"`
	Email       string `gorm:"size:64"`
	AvatarURL   string `gorm:"size:255"`
}

func (u UserExtra) TableName() string {
	return "user_extra"
}
