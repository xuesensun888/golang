package model

import (
	"fmt"
	"remotehttp/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Instance struct {
	DB *gorm.DB
}

var Conn *Instance

type AppId struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
type AppVersion struct {
	ID        string `json:"id"`
	AppId     string `json:"app_id"`
	Version   string `json:"version"`
	Md5Sum    string `gorm:"column:md5" json:"md5"`
	Path      string `gorm:"column:url" json:"url"`
	TimeStamp string `gorm:"column:date" json:"date"`
	Type      string `gorm:"type"`
}

func (i *Instance) AppId(app string) string {
	var id AppId
	res := i.DB.Table("io_app").Where("name = ?", app).First(&id)
	if res.Error != nil {
		panic(fmt.Errorf("select first line error: %s", res.Error))
	}
	return id.ID
}
func (i *Instance) GetVersion(app string) (v *AppVersion, err error) {
	id := i.AppId(app)
	var ver AppVersion
	res := i.DB.Table("io_history").Where("app_id = ?", id).Order("date desc").First(&ver)
	if res.Error != nil {
		err = res.Error
		return
	}
	v = &ver
	return
}
func NewConn() *Instance {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", conf.DefaultConfig.User, conf.DefaultConfig.Password,
		conf.DefaultConfig.Host, conf.DefaultConfig.Port, conf.DefaultConfig.DbName, conf.DefaultConfig.CharSet,
		conf.DefaultConfig.ParseTime, conf.DefaultConfig.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Instance{
		DB: db,
	}

}
