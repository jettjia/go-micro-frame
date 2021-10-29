package oss

import (
	"fmt"
	"os"
	"testing"
)

// 指定文件路径上传
func Test_Aliyun_UploadByFilepath(t *testing.T) {
	Aliyun.AccessKeyId = "xx"
	Aliyun.AccessKeySecret = "xx"
	Aliyun.Endpoint = "oss-cn-beijing.aliyuncs.com"
	Aliyun.BucketName = "bucket-name"
	Aliyun.BasePath = "test"
	Aliyun.Filename = "xxx.png"

	path, filename, err := Aliyun.UploadByFilepath("test.png")

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}

// 直接上传文件
func Test_Aliyun_UploadByFile(t *testing.T) {
	Aliyun.AccessKeyId = "xx"
	Aliyun.AccessKeySecret = "xx"
	Aliyun.Endpoint = "oss-cn-beijing.aliyuncs.com"
	Aliyun.BucketName = "xx"
	Aliyun.BasePath = "test"
	Aliyun.Filename = "xxx.png"

	var file *os.File
	file, err := os.Open("test.png")
	if err != nil {
		t.Error(err.Error())
	}

	path, filename, err := Aliyun.UploadByFile(file)

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}