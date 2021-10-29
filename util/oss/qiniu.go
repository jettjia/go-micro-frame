package oss

import (
	"context"
	"fmt"
	"github.com/jettjia/go-micro-frame/util/oss/interfaces"
	"mime/multipart"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

var _ interfaces.Oss = (*qiniu)(nil)

var Qiniu = new(qiniu)

type qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速

	Filename string
	Filesize int64
}

func (q *qiniu) DeleteByKey(key string) error {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	config := q.GetConfig()
	bucketManager := storage.NewBucketManager(mac, config)
	if err := bucketManager.Delete(q.Bucket, key); err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}
	return nil
}

func (q *qiniu) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	var result storage.PutRet

	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	putPolicy := storage.PutPolicy{Scope: q.Bucket}
	uploadToken := putPolicy.UploadToken(mac)

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	formUploader := storage.NewFormUploader(q.GetConfig())
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}
	err = formUploader.Put(context.Background(), &result, uploadToken, q.Filename, file, q.Filesize, &putExtra)
	if err != nil {
		return filepath, filename, err
	}

	filename = result.Key
	filepath = q.ImgPath + "/" + filename
	return filepath, filename, nil
}

func (q *qiniu) UploadByFilepath(p string) (path string, filename string, err error) {
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

	q.Filesize = info.Size()

	//_, q.Filename = filepath.Split(path)
	return q.UploadByFile(file)
}

func (q *qiniu) UploadByFileHeader(file *multipart.FileHeader) (filepath string, filename string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filepath, filename, err
	}
	q.Filename = fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	q.Filesize = file.Size
	return q.UploadByFile(open)
}

func (q *qiniu) GetConfig() *storage.Config {
	cfg := storage.Config{
		Zone:          q.GetZone(),
		UseHTTPS:      q.UseHTTPS,
		UseCdnDomains: q.UseCdnDomains,
	}
	return &cfg
}

func (q *qiniu) GetZone() *storage.Region {
	switch q.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuaDong":
		return &storage.ZoneHuadong
	case "ZoneHuaBei":
		return &storage.ZoneHuabei
	case "ZoneHuaNan":
		return &storage.ZoneHuanan
	case "ZoneBeiMei":
		return &storage.ZoneBeimei
	case "ZoneXinJiaPo":
		return &storage.ZoneXinjiapo
	}
	return &storage.ZoneHuadong
}
