package conf

import (
	"github.com/spf13/viper"
	"log"
)

var Cfg *viper.Viper

func init() {
	Cfg = viper.New()
	Cfg.SetConfigName("conf")
	Cfg.SetConfigType("json")
	Cfg.AddConfigPath("./conf")
	if err := Cfg.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件失败,", err)
		return
	}
	initMySql()
}
