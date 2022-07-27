package core

import (
	"errors"
	"typora-qiniu-upload/utils"
)

/**
 *  @Description: 上传图片
 *  @Author: HanPengYu
 *  @param rootFlag
 *  @param args
 *  @return string
 *  @return error
 **/
func UploadImg(fileList []string) (string, error) {
	var rstStr string = "Upload Faild:"

	var succFiles = make([]string, 0)
	for _, fileInfo := range fileList {
		imageUrl, err := utils.UploadImageTuQiNiuByForm(fileInfo)
		if err == nil {
			succFiles = append(succFiles, imageUrl)
		}
	}

	if len(succFiles) == 0 {
		return utils.ErrorString(rstStr, "文件全部上传失败"), errors.New("文件全部上传失败")
	}

	return utils.SuccessString("Upload Success:", succFiles), nil
}
