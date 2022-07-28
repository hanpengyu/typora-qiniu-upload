package cmd

import (
	"github.com/spf13/pflag"
	"typora-qiniu-upload/common/config"
	logger "typora-qiniu-upload/common/log"
)

/**
 *  @Description: 启动加载相关
 *  @Author: HanPengYu
 *  @param rootFlag
 **/
func bootstrap(rootFlag *pflag.FlagSet) {
	// 初始化配置
	config.InitConfig(rootFlag)

	// 日志初始化
	logger.InitLog()

	last()
}

func last() {
	defer logger.GetLog().Sync()
}
