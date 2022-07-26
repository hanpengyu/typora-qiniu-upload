package utils

import (
	"github.com/spf13/viper"
	"typora-qiniu-upload/config"
)

func Config() *viper.Viper {
	return config.GetConfig()
}

func GetConfigString(key string) string {
	return config.GetConfig().GetString(key)
}
