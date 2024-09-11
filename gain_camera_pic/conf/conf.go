package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	User         string
	Passwd       string
	Host         string
	Port         string
	DbName       string
	Charset      string
	ParseTime    string
	Loc          string
	DownloadPath string
}

var DefaultConfig Config

func LoadConfig(conf string) {
	viper.SetConfigFile(conf)
	viper.SetConfigType("json")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	DefaultConfig.User = viper.GetString("user")
	DefaultConfig.Passwd = viper.GetString("passwd")
	DefaultConfig.Host = viper.GetString("host")
	DefaultConfig.Port = viper.GetString("port")
	DefaultConfig.DbName = viper.GetString("dbname")
	DefaultConfig.Charset = viper.GetString("charset")
	DefaultConfig.ParseTime = viper.GetString("parsetime")
	DefaultConfig.Loc = viper.GetString("loc")
	DefaultConfig.DownloadPath = viper.GetString("downloadpath")

}
