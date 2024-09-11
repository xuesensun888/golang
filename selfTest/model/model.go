package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int64
	Name string
}
type Instance struct {
	DB *gorm.DB
}

var Conn *Instance

func (i *Instance) User(Id int64) User {
	var user User
	i.DB.Table("user").Where("id = ?", Id).First(&user)
	return user

}
func NewConn() *Instance {
	dsn := "root:Cargo1e9@2024@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Instance{
		DB: db,
	}

}
