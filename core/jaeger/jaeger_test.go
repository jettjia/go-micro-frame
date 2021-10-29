package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"testing"
	"time"
)

func Test_InitJaeger(t *testing.T) {
	// 初始化
	InitJaeger("10.4.7.71", 6831, "gomicrom-test", "", "")

	// 上传
	span := opentracing.StartSpan("gomicrom.com.test")
	time.Sleep(time.Second)
	defer span.Finish()
}