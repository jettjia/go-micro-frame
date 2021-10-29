package util

import (
	"crypto/md5"
	"encoding/hex"
)

var Encrypt = new(encrypt)

type encrypt struct {}

func (e *encrypt) Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
