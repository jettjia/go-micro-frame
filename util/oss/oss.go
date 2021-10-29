package oss

import "github.com/jettjia/go-micro-frame/util/oss/interfaces"

func Oss(ossType string) interfaces.Oss {
	switch ossType {
	case "local":
		return Local
	case "qiniu":
		return Qiniu
	case "aliyun":
		return Aliyun
	//case "tencent":
	//	return Tencent
	default:
		return Qiniu
	}
}
