package interfaces

import (
	"mime/multipart"
)

type Oss interface {
	// DeleteByKey 通过 key(唯一标识) 删除在oss上的文件
	DeleteByKey(key string) error
	// UploadByFile 方便使用 os.File 数据上传oss
	UploadByFile(file multipart.File) (filepath string, filename string, err error)
	// UploadByFilepath 通过文件路径上传文件到oss
	UploadByFilepath(path string) (filepath string, filename string, err error)
	// UploadByFileHeader 方便从http框架接收到的 multipart.FileHeader 数据上传oss
	UploadByFileHeader(header *multipart.FileHeader) (filepath string, filename string, err error)
}
