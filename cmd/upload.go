package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"typora-qiniu-upload/utils"
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

/**
 *  @Description: 上传图片
 *  @Author: HanPengYu
 *  @param args
 **/
func upload(args []string) {
	var rstStr string = "Upload Faild:"

	var succFiles = make([]string, 0)
	for _, fileInfo := range args {
		imageUrl, err := utils.UploadImageTuQiNiuByForm(fileInfo)
		if err == nil {
			succFiles = append(succFiles, imageUrl)
		}
	}

	if len(succFiles) == 0 {
		fmt.Println(utils.ErrorString(rstStr, "文件全部上传失败"))
		return
	}

	fmt.Println(utils.SuccessString("Upload Success:", succFiles))
}
