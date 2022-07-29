package utils

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
	"typora-qiniu-upload/common/config"
	logger "typora-qiniu-upload/common/log"
)

/**
 *  @Description: 图片压缩,返回压缩之后的本地图片路径
 *  @Author: HanPengYu
 *  @return string
 *  @return error
 **/
func ImageCompress(filePath string) string {
	compressCfg := config.GetImageCompressCfg()

	// 未开始压缩图片设置
	if compressCfg.CompressSwitch == 0 {
		return filePath
	}

	file, err := os.Open(filePath)
	if err != nil {
		logger.Info("打开文件失败", filePath)
		return ""
	}
	defer file.Close()

	// 判断文件大小
	statInfo, err := file.Stat()
	if err != nil {
		return ""
	}

	logger.Info("stat info", statInfo)
	if statInfo.Size() <= int64(compressCfg.MaxKb*1000) {
		logger.Info("文件太小不用压缩", filePath)
		return filePath
	}

	ext := filepath.Ext(filePath)
	ext = strings.Trim(ext, ".")
	allowExtMap := map[string]struct{}{
		"jpg":  struct{}{},
		"jpeg": struct{}{},
		"png":  struct{}{},
	}
	if _, ok := allowExtMap[ext]; !ok {
		return filePath
	}

	imageConfig, imageDecode, err := imageInfo(file, ext)
	if err != nil {
		return ""
	}

	// 等比压缩尺寸
	compressWidth := compressCfg.Width
	width := uint(compressWidth)
	height := uint(compressWidth * imageConfig.Height / imageConfig.Width)

	//  创建存放压缩之后图片的文件夹
	t := time.Now()
	ymd := t.Format("20060102")
	hms := t.Format("150405")
	tmpDir := "/tmp/" + ymd
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err := os.Mkdir(tmpDir, 0755)
		if err != nil {
			panic("创建目录失败" + tmpDir)
		}
	}

	// 压缩之后写入新文件的路径 | /tmp/20220727/c_width×height_hms_name
	newFileName := fmt.Sprintf("%s/c_%d×%d_%s_%s", tmpDir, width, height, hms, statInfo.Name())
	logger.Info("压缩后图片名字", newFileName)
	newFile, err := os.Create(newFileName)
	if err != nil {
		return ""
	}
	defer newFile.Close()

	// 执行图片压缩
	newImage := resize.Thumbnail(width, height, imageDecode, resize.Lanczos3)
	err = imageWrite(newImage, newFile, ext)
	if err != nil {
		return ""
	}

	return newFileName
}

/**
 *  @Description: 图片信息
 *  @Author: HanPengYu
 *  @param file
 *  @param ext
 *  @return image.Config
 *  @return image.Image
 *  @return error
 **/
func imageInfo(file *os.File, ext string) (image.Config, image.Image, error) {
	var imageConfig image.Config
	var imageDecode image.Image
	var err error

	switch {
	case ext == "jpg" || ext == "jpeg":
		imageConfig, err = jpeg.DecodeConfig(file)
		if err != nil {
			return imageConfig, imageDecode, err
		}

		_, _ = file.Seek(0, 0)
		imageDecode, err = jpeg.Decode(file)
		if err != nil {
			return imageConfig, imageDecode, err
		}
	case ext == "png":
		imageConfig, err = png.DecodeConfig(file)
		if err != nil {
			return imageConfig, imageDecode, err
		}

		_, _ = file.Seek(0, 0)
		imageDecode, err = png.Decode(file)
		if err != nil {
			return imageConfig, imageDecode, err
		}
	}

	return imageConfig, imageDecode, nil
}

/**
 *  @Description: 把压缩后的图片写入文件中
 *  @Author: HanPengYu
 *  @param imageFile
 *  @param outFile
 *  @param ext
 *  @return error
 **/
func imageWrite(imageFile image.Image, outFile *os.File, ext string) error {
	var err error
	switch {
	case ext == "jpg" || ext == "jpeg":
		err = jpeg.Encode(outFile, imageFile, &jpeg.Options{Quality: 80})
	case ext == "png":
		err = png.Encode(outFile, imageFile)
	}
	return err
}
