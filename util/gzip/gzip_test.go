package gzip

import "testing"

// 压缩
func TestCompress(t *testing.T) {
	zipClient := NewGzip()
	err := zipClient.Compress("tmp", "tmp.zip")

	if err != nil {
		t.Error(err.Error())
	}
}

// 解压
func TestDeCompress(t *testing.T) {
	zipClient := NewGzip()
	// 为空表示解压到当前目录
	zipFileName := "tmp.zip"
	err := zipClient.DeCompress( zipFileName)
	if err != nil {
		t.Error(err.Error())
	}
}
