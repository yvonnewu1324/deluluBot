package helper

import (
	"github.com/spf13/viper"
)

func NewViper() {
	viper.SetConfigName("gowatch")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil && viper.GetString("ENV") != "PROD" {
		LogError("fatal error config file: %w", err)
	}
}
