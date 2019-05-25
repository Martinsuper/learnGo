package aMoveImage

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"log"
	"os"
	"path"
	"path/filepath"
)

func HandleImage(projectpath, saveImagePath string) {
	ans := FindImage(projectpath, ".")
	if !ans {
		log.Fatal("该文件夹下不存在图片")
		return
	}
	absPath, _ := filepath.Abs("./")
	res := GetImagePath(absPath + "/imagelist.json")
	for i := 0; i < len(res); i++ {
		handleAimage(res[i], saveImagePath)
	}
}
func handleAimage(imageName, saveImagePath string) {
	//获取文件名
	filenameWithSuffix := path.Base(imageName)
	//读取文件头部信息并创建文件夹，返回文件夹路径
	projPath := CreateProject(ReadImage(imageName), saveImagePath)
	//移动文件到指定文件夹下
	MoveFile(imageName, projPath+"/"+filenameWithSuffix)
}
func MoveFile(oldpath, newpath string) {
	err := os.Rename(oldpath, newpath)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadImage(imageName string) string {
	f, err := os.Open(imageName)
	if err != nil {
		log.Fatal(err)
	}
	exif.RegisterParsers(mknote.All...)
	x, err1 := exif.Decode(f)
	if err1 != nil {
		return ""
	}
	tm, err2 := x.DateTime()
	if err2 != nil {
		fmt.Println(err2)
		log.Fatal(err2)
	}
	return tm.String()
}
