package methodset

import "hash/crc32"

// GenUniqueStringToInt 传入字符串，生成唯一的int值
func GenUniqueStringToInt(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
