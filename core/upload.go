package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/pflag"
	"path/filepath"
	"time"
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
func UploadImg(rootFlag *pflag.FlagSet, args []string) (string, error) {
	var rstStr string = "Upload Faild"

	// 配置获取
	accessKey := utils.GetConfigString("QiNiu.AccessKey")
	secretKey := utils.GetConfigString("QiNiu.SecretKey")
	bucket := utils.GetConfigString("QiNiu.Bucket")
	prefix := utils.GetConfigString("QiNiu.Prefix")
	doman := utils.GetConfigString("QiNiu.Domain")

	// 上传配置类
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	// 上传策略
	putPolicy := storage.PutPolicy{Scope: bucket}

	// 生成上传凭证
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 表单上传对象
	formUploader := storage.NewFormUploader(&cfg)

	var succFiles = make([]string, 0)
	for _, fileInfo := range args {
		// todo 压缩图片

		// 文件名
		t := time.Now().Format("200601021504_05")
		_, fName := filepath.Split(fileInfo)
		fileName := fmt.Sprintf("%s_%s", t, fName)

		// 表单上传
		ret := storage.PutRet{}
		putExtra := storage.PutExtra{}
		key := fmt.Sprintf("%s/%s", prefix, fileName)
		err := formUploader.PutFile(context.Background(), &ret, upToken, key, fileInfo, &putExtra)
		if err == nil {
			f := fmt.Sprintf("%s/%s", doman, key)
			succFiles = append(succFiles, f)
			// todo 删掉原图
		}
	}

	if len(succFiles) == 0 {
		return utils.ErrorString(rstStr, "文件全部上传失败"), errors.New("文件全部上传失败")
	}

	return utils.SuccessString("Upload Success:", succFiles), nil
}
