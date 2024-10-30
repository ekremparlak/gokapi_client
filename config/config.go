package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.ReadInConfig()
}
