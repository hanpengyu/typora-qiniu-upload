package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"hello-qiniu/core"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "typora 上传图片到七牛云",
		Run: func(cmd *cobra.Command, args []string) {
			upload(rootFlag, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(uploadCmd)
}

func upload(rootFlag *pflag.FlagSet, args []string) {
	rstStr, err := core.UploadImg(rootFlag, args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rstStr)
}
