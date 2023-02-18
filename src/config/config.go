package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() (err error) {

	viper.AddConfigPath("/pong-app/")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	return
}
