package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"path/filepath"
	"time"
)

/**
 *  @Description: 上传文件到七牛云中
 *  @Author: HanPengYu
 *  @param filePath
 *  @return string
 *  @return error
 **/
func UploadImageTuQiNiuByForm(filePath string) (string, error) {
	// 配置获取
	accessKey := GetConfigString("QiNiu.AccessKey")
	secretKey := GetConfigString("QiNiu.SecretKey")
	bucket := GetConfigString("QiNiu.Bucket")
	prefix := GetConfigString("QiNiu.Prefix")
	doman := GetConfigString("QiNiu.Domain")

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

	// 压缩图片
	filePath = ImageCompress(filePath)

	t := time.Now().Format("200601021504_05")
	_, fName := filepath.Split(filePath)
	fileName := fmt.Sprintf("%s_%s", t, fName)

	// 表单上传
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	key := fmt.Sprintf("%s/%s", prefix, fileName)
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", doman, key), nil
}
