package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var conf *Config

type Config struct {
	cfg *viper.Viper
}

func GetConfig() *viper.Viper {
	return conf.cfg
}

func setConfig(c *Config) {
	conf = c
}

func InitConfig(f *pflag.FlagSet) {
	rootViper := viper.New()
	err := rootViper.BindPFlags(f)
	if err != nil {
		panic("配置文件绑定错误")
	}

	// 配置文件路径
	configPath := rootViper.GetString("config")
	if configPath == "" {
		configPath = "/usr/local/etc/qiniu.toml"
	}

	// 解析配置文件
	qnViper := viper.New()
	qnViper.SetConfigFile(configPath)
	err = qnViper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败:%s", configPath))
	}

	conf := &Config{
		cfg: qnViper,
	}
	setConfig(conf)
}

func InitConfigTest() {
	configPath := "/usr/local/etc/qiniu.toml"
	// 解析配置文件
	qnViper := viper.New()
	qnViper.SetConfigFile(configPath)
	err := qnViper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败:%s", configPath))
	}

	conf := &Config{
		cfg: qnViper,
	}
	setConfig(conf)
}
