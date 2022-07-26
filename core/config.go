package core

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitConfig(f *pflag.FlagSet) (*viper.Viper, error) {
	rootViper := viper.New()
	err := rootViper.BindPFlags(f)
	if err != nil {
		return nil, errors.New("配置文件绑定错误")
	}

	// 配置文件路径
	configPath := rootViper.GetString("config")
	if configPath == "" {
		configPath = "./config/qiniu.toml"
	}

	// 解析配置文件
	qnViper := viper.New()
	qnViper.SetConfigFile(configPath)
	err = qnViper.ReadInConfig()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("配置文件解析失败:%s", configPath))
	}

	return qnViper, nil
}
