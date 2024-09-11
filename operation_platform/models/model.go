package models

import (
	"fmt"
	"operation/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Instance struct {
	DB *gorm.DB
}

var Conn *Instance

func NewConn() *Instance {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", conf.DefaultConfig.User, conf.DefaultConfig.Password, conf.DefaultConfig.Host,
		conf.DefaultConfig.Port, conf.DefaultConfig.Dbname, conf.DefaultConfig.CharSet, conf.DefaultConfig.ParseTime, conf.DefaultConfig.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&FaultReport{})
	return &Instance{
		DB: db,
	}

}
