package oss

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	util "github.com/jettjia/go-micro-frame/util/encrypt"
	"github.com/jettjia/go-micro-frame/util/oss/interfaces"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var _ interfaces.Oss = (*local)(nil)

var Local = new(local)

type local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径

	Filename string
	Filesize int64
}

func (l *local) DeleteByKey(key string) error {
	path := l.Path + "/" + key
	if strings.Contains(path, l.Path) {
		if err := os.Remove(path); err != nil {
			return errors.Wrap(err, "本地文件删除失败!")
		}
	}
	return nil
}

func (l *local) UploadByFile(file multipart.File) (filepath string, Filename string, err error) {
	filepath = l.CreateFilepath(l.Filename)
	var out *os.File
	if out, err = os.Create(filepath); err != nil {
		return filepath, Filename, errors.Wrap(err, "读取文件失败!")
	}

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("open 文件关闭失败!", zap.Error(err))
		}
		if err = out.Close(); err != nil {
			zap.L().Error("out 文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	if _, err = io.Copy(out, file); err != nil {
		return filepath, Filename, errors.Wrap(err, "传输(拷贝)文件失败!")
	} // 传输(拷贝)文件
	return filepath, Filename, nil
}

func (l *local) UploadByFilepath(p string) (path string, Filename string, err error) {
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
	l.Filesize = info.Size()
	//_, l.Filename = filepath.Split(path)
	return l.UploadByFile(file)
}

func (l *local) UploadByFileHeader(file *multipart.FileHeader) (filepath string, Filename string, err error) {
	if err = os.MkdirAll(l.Path, os.ModePerm); err != nil {
		return filepath, Filename, errors.Wrap(err, "创建路径失败!")
	} // 尝试创建此路径
	l.Filename = l.CreateFilename(file.Filename)
	var open multipart.File
	if open, err = file.Open(); err != nil {
		return filepath, Filename, errors.Wrap(err, "读取文件失败!")
	} // 读取文件
	return l.UploadByFile(open)
}

// Filename 拼接新文件名
// Author [SliverHorn](https://github.com/SliverHorn)
func (l *local) CreateFilename(Filename string) string {
	ext := path.Ext(Filename)                 // 读取文件后缀
	name := strings.TrimSuffix(Filename, ext) // 读取文件名并加密
	Filename = util.Encrypt.Md5([]byte(name))
	return Filename + "_" + time.Now().Format("2006_01_02_15_04_05") + ext
}

// Filepath 拼接路径和文件名
// Author [SliverHorn](https://github.com/SliverHorn)
func (l *local) CreateFilepath(Filename string) string {
	return l.Path + "/" + Filename
}
