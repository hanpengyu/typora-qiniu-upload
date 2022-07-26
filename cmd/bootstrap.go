package cmd

import (
	"github.com/spf13/pflag"
	"typora-qiniu-upload/config"
)

/**
 *  @Description: 启动加载相关
 *  @Author: HanPengYu
 *  @param rootFlag
 **/
func bootstrap(rootFlag *pflag.FlagSet) {
	// 初始化配置
	config.InitConfig(rootFlag)
}
