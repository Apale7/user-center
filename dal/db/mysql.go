package db

import (
	"fmt"
	"user_center/model"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	db_model "user_center/dal/db/model"
)

const (
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	base = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//获取dbconf
	mysqlConf := getDbConf()
	//构造dsn
	dsn := fmt.Sprintf(base, mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	//连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})//关闭默认事务提高性能
	if err != nil {
		panic(err)
	}
	log.Info("mysql连接成功")
	//建表
	if !db.Migrator().HasTable(&db_model.User{}) {
		db.Migrator().CreateTable(&db_model.User{})
	}
	if !db.Migrator().HasTable(&db_model.UserExtra{}) {
		db.Migrator().CreateTable(&db_model.UserExtra{})
	}
}

func getDbConf() model.Mysql {
	viper.SetConfigName("db_conf")
	viper.AddConfigPath("../../config")
	if err = viper.ReadInConfig(); err != nil {
		log.Error(errors.WithStack(err))
		panic("viper readInConfig error")
	}
	var dbconf model.Conf
	if err = viper.Unmarshal(&dbconf); err != nil {
		log.Error(errors.WithStack(err))
		panic("viper Unmarshal error")
	}
	return dbconf.Mysql
}

func GetDb() *gorm.DB {
	return db
}
