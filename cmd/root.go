package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	rootCmd = &cobra.Command{
		Use:   "TyporaQiNiuUpload",
		Short: "七牛上传文件入口",
	}
	cfgFile  string
	rootFlag *pflag.FlagSet
)

func init() {
	rootFlag = rootCmd.PersistentFlags()
	rootFlag.StringVarP(&cfgFile, "config", "c", "/usr/local/etc/qiniu.toml", "配置文件")

	// 启动相关
	bootstrap(rootFlag)
}

func Execute() {
	_ = rootCmd.Execute()
}
