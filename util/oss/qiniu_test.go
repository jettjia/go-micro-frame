package oss

import (
	"fmt"
	"os"
	"testing"
)

func Test_Qiniu_UploadByFilepath(t *testing.T) {
	Qiniu.AccessKey = "ww"
	Qiniu.SecretKey = "ww"
	Qiniu.Bucket = "test-zhebaihui"
	Qiniu.UseHTTPS = false
	Qiniu.Filename = "xxx.png"
	Qiniu.ImgPath = "http://ww.com"

	path, filename, err := Qiniu.UploadByFilepath("./test.png")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}

// 直接上传文件
func Test_Qiniu_UploadByFile(t *testing.T) {
	Qiniu.AccessKey = "xx"
	Qiniu.SecretKey = "xx"
	Qiniu.Bucket = "test-zhebaihui"
	Qiniu.UseHTTPS = false
	Qiniu.Filename = "xxx.png"
	Qiniu.ImgPath = "http://img.xx.com"



	var file *os.File
	file, err := os.Open("test.png")
	if err != nil {
		t.Error(err.Error())
	}

	// 需要设置大小
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		t.Error(err.Error())
	}
	Qiniu.Filesize = info.Size()

	path, filename, err := Qiniu.UploadByFile(file)

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}
