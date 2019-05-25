package main

import (
	"flag"
	"mAlgorithm/aMoveImage"
)

func main() {
	var findImagePath string
	var saveImagePath string
	flag.StringVar(&findImagePath, "f", "./", "输入要移动图片的目录")
	flag.StringVar(&saveImagePath, "t", "./", "输入要保存图片的路径")
	flag.Parse()
	aMoveImage.HandleImage(findImagePath, saveImagePath)
}
