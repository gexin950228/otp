package dataSource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"otp/logSource"
)

var Db *gorm.DB
var err error

func InitDb() *gorm.DB {
	dataSource := LoadMysqlConfig()
	host := dataSource.Host
	port := dataSource.Port
	user := dataSource.User
	password := dataSource.Password
	database := dataSource.DataBase
	logMode := dataSource.LogMode
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, database)
	logrus.Info(fmt.Sprintf("Mysql connections: %s", dst))
	Db, err = gorm.Open("mysql", dst)
	Db.LogMode(logMode)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	logSource.Log.Info("数据库初始化成功！")
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(10)
	//Db.AutoMigrate(&models.Department{}, &models.Machine{}, &models.VPNInfo{})
	return Db
}
