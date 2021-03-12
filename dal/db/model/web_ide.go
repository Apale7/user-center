package model

import (
	"gorm.io/gorm"
)

// Container [...]
type Container struct {
	gorm.Model
	ContainerID string `gorm:"unique;column:container_id;type:varchar(80);not null"`         // docker容器id
	Status      int8   `gorm:"column:status;type:tinyint(4);not null"`                       // Docker容器状态 1-Running 2-Paused 3-Restarting 4-OOMKilled 5-Dead 6-UNKNOWN
	ImageID     string `gorm:"index:idx_image_id;column:image_id;type:varchar(80);not null"` // docker镜像id
	Name        string `gorm:"column:name;type:varchar(32);not null"`                        // docker容器名字
}

// Group [...]
type Group struct {
	gorm.Model
	OwnerID   uint32 `gorm:"uniqueIndex:uniq_idx_owner_id_group_name;column:owner_id;type:int(11) unsigned;not null"`          // 组长的user_id
	GroupName string `gorm:"uniqueIndex:uniq_idx_owner_id_group_name;column:group_name;type:varchar(32);not null;default:未命名"` // 组名
}

// Image [...]
type Image struct {
	gorm.Model
	ImageID   string `gorm:"unique;column:image_id;type:varchar(80);not null"` // docker镜像id
	ImageSize int64  `gorm:"column:image_size;type:bigint(20);not null"`       // 镜像大小，单位byte
	Author    string `gorm:"column:author;type:varchar(32);not null"`          // 镜像作者
	RepoTags  string `gorm:"column:repo_tags;type:text"`                       // 仓库标签，以json数组的形式存储
}

// Resource [...]
type Resource struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(16);not null"` // 资源名称
}

// Role [...]
type Role struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(16);not null"` // 角色名称
}

// RoleResource [...]
type RoleResource struct {
	gorm.Model
	RoleID     uint32 `gorm:"index:idx_user_id;column:role_id;type:int(11) unsigned;not null"` // role表的id
	ResourceID uint32 `gorm:"column:resource_id;type:int(11) unsigned;not null"`               // resource表的id
}

// User [...]
type User struct {
	gorm.Model
	Username string `gorm:"unique;unique;column:username;type:varchar(16);not null"` // 用户名
	Password string `gorm:"column:password;type:varchar(32);not null"`               // 密码
}

// UserContainer [...]
type UserContainer struct {
	gorm.Model
	UserID      uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	ContainerID string `gorm:"column:container_id;type:varchar(80);not null"`                   // docker中的container_id
}

// UserExtra [...]
type UserExtra struct {
	gorm.Model
	UserID      uint32 `gorm:"unique;column:user_id;type:int(11) unsigned;not null"` // user表的id
	Nickname    string `gorm:"column:nickname;type:varchar(16);not null"`            // 昵称
	PhoneNumber string `gorm:"column:phone_number;type:varchar(16)"`                 // 电话号码
	Email       string `gorm:"column:email;type:varchar(32)"`                        // 邮箱
	AvatarURL   string `gorm:"column:avatar_url;type:varchar(255)"`                  // 头像链接
}

// UserGroup [...]
type UserGroup struct {
	gorm.Model
	UserID  uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	GroupID string `gorm:"index:idx_group_id;column:group_id;type:varchar(80);not null"`    // docker中的image_id
}

// UserImage [...]
type UserImage struct {
	gorm.Model
	UserID  uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	ImageID string `gorm:"column:image_id;type:varchar(80);not null"`                       // docker中的image_id
}

// UserRole [...]
type UserRole struct {
	gorm.Model
	UserID uint32 `gorm:"index:idx_user_id;column:user_id;type:int(11) unsigned;not null"` // user表的id
	RoleID uint32 `gorm:"column:role_id;type:int(11) unsigned;not null"`                   // role表的id
}
