package logger

import (
	"errors"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	NewLogger("go-micro-frame", "./logs/go-micro-frame-test.log", "", 128, 30,7)
	err := errors.New("i am test log info")
	Error(err)
}