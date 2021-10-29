package oss

import (
	"fmt"
	"os"
	"testing"
)

// 直接上传文件
func Test_Local_UploadByFile(t *testing.T) {
	Local.Path = "/tmp" //上传保存的路径
	Local.Filename = "test.png" // 上传的文件的名字

	var file *os.File
	file, err := os.Open("test.png")
	if err != nil {
		t.Error(err.Error())
	}

	path, filename, err := Local.UploadByFile(file)

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}

// 通过文件路径上传文件到oss
func Test_Local_UploadByFilepath(t *testing.T) {
	Local.Path = "/tmp" //上传保存的路径

	Local.Filename = "test.png" // 上传的文件的名字

	path, filename, err := Local.UploadByFilepath("test.png")

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}