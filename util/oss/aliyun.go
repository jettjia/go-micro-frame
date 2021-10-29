package oss

import (
	"fmt"
	"github.com/jettjia/go-micro-frame/util/oss/interfaces"
	"mime/multipart"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

var _ interfaces.Oss = (*aliyun)(nil)

var Aliyun = new(aliyun)

type aliyun struct {
	BasePath        string `mapstructure:"base-path" json:"basePath" yaml:"base-path"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`

	Filename string
	Filesize int64
}

func (a *aliyun) DeleteByKey(key string) error {
	bucket, err := a.getAliyunBucket()
	if err != nil {
		return err
	}

	if err = bucket.DeleteObject(key); err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}

	return nil
}

func (a *aliyun) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	bucket, newErr := a.getAliyunBucket()
	if newErr != nil {
		return filepath, filename, newErr
	}

	defer func() {
		_ = file.Close()
	}() // 关闭文件流

	filepath = a.Filepath(a.Filename)
	err = bucket.PutObject(filepath, file) // 上传文件流。
	if err != nil {
		fmt.Println("上传阿里云oss, err", err.Error())
		return "", "", err
	}
	return filepath, a.Filename, nil
}

func (a *aliyun) UploadByFilepath(p string) (path string, filename string, err error) {
	var file *os.File
	file, err = os.Open(p)
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件不存在!", p)
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件信息获取失败!", p)
	}
	a.Filesize = info.Size()

	//_, a.Filename = filepath.Split(path)
	return a.UploadByFile(file)
}

func (a *aliyun) UploadByFileHeader(file *multipart.FileHeader) (filepath string, filename string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filepath, filename, err
	}
	a.Filename = file.Filename
	return a.UploadByFile(open)
}

// Filepath 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *aliyun) Filepath(Filename string) string {
	return a.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + Filename
}

func (a *aliyun) getAliyunBucket() (bucket *oss.Bucket, err error) {
	var client *oss.Client
	if client, err = oss.New(a.Endpoint, a.AccessKeyId, a.AccessKeySecret); err != nil {
		return nil, err
	} // 创建OSSClient实例

	if bucket, err = client.Bucket(a.BucketName); err != nil {
		return nil, errors.Wrap(err, "获取存储空间失败!")
	} // 获取存储空间

	return bucket, nil
}
