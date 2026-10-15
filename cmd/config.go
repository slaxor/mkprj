package main

import (
	"log"

	"github.com/spf13/viper"
)

var (
	config  *viper.Viper
	appname = "mkprj"
)

func init() {
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME/.config/" + appname)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
