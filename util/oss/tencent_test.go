package oss

import (
	"fmt"
	"os"
	"testing"
)

func Test_UploadByFile(t *testing.T) {
	Tencent.Bucket="xx"
	Tencent.SecretID="xx"
	Tencent.SecretKey="xx"
	Tencent.Region="xx"
	Tencent.BaseURL="xx"
	Tencent.PathPrefix="xx"

	var file *os.File
	file, err := os.Open("test.png")
	if err != nil {
		t.Error(err.Error())
	}

	path, filename, err := Tencent.UploadByFile(file)

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("------------path:", path)
	fmt.Println("------------Filename:", filename)
}
