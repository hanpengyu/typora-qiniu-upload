package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"path/filepath"
	"strings"
	"time"
	"typora-qiniu-upload/common/config"
	logger "typora-qiniu-upload/common/log"
)

/**
 *  @Description: 上传文件到七牛云中
 *  @Author: HanPengYu
 *  @param filePath
 *  @return string
 *  @return error
 **/
func UploadImageTuQiNiuByForm(filePath string) (string, error) {
	logger.Info("源文件", filePath)

	// 配置获取
	qnCfg := config.GetQnCfg()

	// 上传配置类
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	// 上传策略
	putPolicy := storage.PutPolicy{Scope: qnCfg.Bucket}

	// 生成上传凭证
	mac := qbox.NewMac(qnCfg.AccessKey, qnCfg.SecretKey)
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
	key := fmt.Sprintf("%s/%s", qnCfg.BucketDir, fileName)
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return "", err
	}

	// 如果是截图需要删除掉本地的文件
	if isScreenImage(filePath) {
		_ = os.Remove(filePath)
	}

	return fmt.Sprintf("%s/%s", qnCfg.CdnUrl, key), nil
}

/**
 *  @Description: 判断上传的图片是否是截图
 *  @Author: HanPengYu
 *  @param filepath
 **/
func isScreenImage(fileName string) bool {
	dir := filepath.Dir(fileName)
	if strings.HasSuffix(dir, "Library/Application Support/typora-user-images") {
		return true
	}
	return false
}
