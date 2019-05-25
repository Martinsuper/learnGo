package aMoveImage

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CreateProject(strPath, project string) string {
	if strPath != "" {
		year := strPath[:4]
		month := strPath[5:7]
		path := year + "/" + month
		path = project + "/" + path
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		return path
	}
	return ""
}

func getFileName() string {
	fullFilename := "test/path/test.txt"
	//fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)
	return ""
}

func FindImage(thisPath, listPath string) bool {
	files, _ := ioutil.ReadDir(thisPath)
	flag := false
	for _, f := range files {
		if f.IsDir() {
			FindImage(thisPath+"/"+f.Name(), listPath)
			continue
		}
		fileSuffix := path.Ext(f.Name())
		if imageType(fileSuffix) {
			flag = true
			thisPath, _ := filepath.Abs(thisPath)
			writePath(thisPath+"/"+f.Name(), listPath+"/imagelist.json")
		}
	}
	return flag
}
func imageType(obj string) bool {
	imageTypeList := map[string]int{
		".png": 1,
		".jpg": 1,
	}
	if imageTypeList[obj] == 1 {
		return true
	}
	return false
}
func writePath(data, path string) {
	fileName := path
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = file.WriteString(data + "\n")
}

func GetImagePath(path string) []string {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	slice := make([]string, 100000)
	cnt := 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		slice[cnt] = string(data)
		cnt++
	}
	err1 := os.Remove(path)
	if err1 != nil {
		log.Fatal(err1)
	}
	return slice[:cnt]
}
