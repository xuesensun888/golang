package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	User      string
	Password  string
	Host      string
	Port      string
	Dbname    string
	CharSet   string
	ParseTime string
	Loc       string
}

var DefaultConfig Config

func LoadConfig(conf string) {
	viper.SetConfigFile(conf)
	viper.SetConfigType("json")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	DefaultConfig.User = viper.GetString("user")
	DefaultConfig.Password = viper.GetString("password")
	DefaultConfig.Host = viper.GetString("host")
	DefaultConfig.Port = viper.GetString("port")
	DefaultConfig.Dbname = viper.GetString("dbname")
	DefaultConfig.CharSet = viper.GetString("charset")
	DefaultConfig.ParseTime = viper.GetString("parsetime")
	DefaultConfig.Loc = viper.GetString("loc")

}
