package db

import (
	"fmt"

	config "user_center/config_loader"
	"user_center/dal/db/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	base = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	// 获取dbconf
	dsn := getDSN()
	// 连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务提高性能
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "gormv2_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	log.Info("mysql连接成功")
	db.AutoMigrate(&model.User{}, &model.UserExtra{}, &model.Resource{}, &model.Role{}, &model.RoleResource{}, &model.UserRole{})
}

func getDSN() string {
	username := config.Get("mysql_user")
	password := config.Get("mysql_pass")
	host := config.Get("mysql_host")
	port := config.Get("mysql_port")
	dbname := config.Get("mysql_db")
	dsn := fmt.Sprintf(base, username, password, host, port, dbname)
	return dsn
}
