package oss

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/jettjia/go-micro-frame/util/oss/interfaces"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

var _ interfaces.Oss = (*tencent)(nil)

var Tencent = new(tencent)

type tencent struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`

	Filename string
	Filesize int64
}

func (t *tencent) DeleteByKey(key string) error {
	client, err := t.getTencentClient()
	if err != nil {
		return err
	}
	name := t.PathPrefix + "/" + key
	if _, err = client.Object.Delete(context.Background(), name); err != nil {
		return errors.Wrap(err, "文件删除失败!")
	}
	return nil
}

func (t *tencent) UploadByFile(file multipart.File) (filepath string, Filename string, err error) {
	var client *cos.Client
	client, err = t.getTencentClient()
	if err != nil {
		return filepath, Filename, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	Filename = t.CreateFilename(t.Filename)
	filepath = t.CreateFilepath(Filename)

	_, err = client.Object.Put(context.Background(), Filename, file, nil)
	if err != nil {
		return filepath, t.Filename, errors.Wrap(err, "文件上传失败!")
	}
	return filepath, t.Filename, nil
}

func (t *tencent) UploadByFilepath(p string) (path string, Filename string, err error) {
	var file *os.File
	file, err = os.Open(p)
	if err != nil {
		return path, Filename, errors.Wrapf(err, "(%s)文件不存在!", p)
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return path, Filename, errors.Wrapf(err, "(%s)文件信息获取失败!", p)
	}
	t.Filesize = info.Size()
	_, t.Filename = filepath.Split(path)
	return t.UploadByFile(file)
}

func (t *tencent) UploadByFileHeader(file *multipart.FileHeader) (filepath string, Filename string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filepath, Filename, err
	}
	t.Filename = fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	return t.UploadByFile(open)
}

func (t *tencent) CreateFilename(Filename string) string {
	return t.PathPrefix + "/" + Filename
}

func (t *tencent) CreateFilepath(Filename string) string {
	return t.BaseURL + "/" + Filename
}

func (t *tencent) getTencentClient() (*cos.Client, error) {
	_url, err := url.Parse("https://" + t.Bucket + ".cos." + t.Region + ".myqcloud.com")
	if err != nil {
		return nil, errors.Wrap(err, "url 拼接失败!")
	}
	baseURL := &cos.BaseURL{BucketURL: _url}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	})
	return client, nil
}
