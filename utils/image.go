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
)

const (
	CompressWidth = 400
)

/**
 *  @Description: 图片压缩,返回压缩之后的本地图片路径
 *  @Author: HanPengYu
 *  @return string
 *  @return error
 **/
func ImageCompress(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	// 判断文件大小
	statInfo, err := file.Stat()
	if err != nil {
		return ""
	}

	if statInfo.Size() <= 20000 {
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
	width := uint(CompressWidth)
	height := uint(CompressWidth * imageConfig.Height / imageConfig.Width)

	// 压缩之后写入新文件的路径
	newFileName := fmt.Sprintf("%s/%s", "/tmp", time.Now().Format("200601021504_05_")+statInfo.Name())
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
