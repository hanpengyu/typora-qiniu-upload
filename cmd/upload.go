package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"typora-qiniu-upload/core"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "typora 上传图片到七牛云",
		Run: func(cmd *cobra.Command, args []string) {
			upload(args)
		},
	}
)

func init() {
	rootCmd.AddCommand(uploadCmd)
}

func upload(args []string) {
	rstStr, err := core.UploadImg(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rstStr)
}
