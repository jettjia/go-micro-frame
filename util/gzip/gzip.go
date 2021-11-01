package gzip

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
)

type Gzip struct{}

func NewGzip() *Gzip {
	return &Gzip{}
}

// Compress 压缩
// srcDir 模板文件 如：tmp
// zipFileName 打包成的文件名称 如：dist.zip
func (*Gzip) Compress(srcDir string, zipFileName string) error {
	os.RemoveAll(zipFileName)

	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}

// DeCompress 解压
// zipFileName 压缩文件 如：dist.zip
func (*Gzip) DeCompress(zipFileName string) (err error) {
	// file read
	//打开并读取压缩文件中的内容
	fr, err := zip.OpenReader(zipFileName)
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	//r.reader.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {
		//判断文件该目录文件是否为文件夹
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(file.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		//为文件时，打开文件
		r, err := file.Open()

		//文件为空的时候，打印错误
		if err != nil {
			fmt.Println(err)
			continue
		}
		//这里在控制台输出文件的文件名及路径
		fmt.Println("unzip: ", file.Name)

		//在对应的目录中创建相同的文件
		NewFile, err := os.Create(file.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//将内容复制
		io.Copy(NewFile, r)
		//关闭文件
		NewFile.Close()
		r.Close()
	}

	return nil
}
