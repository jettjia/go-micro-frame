package xlsx

import (
	"fmt"
	"testing"
)

func Test_ImportExcel(t *testing.T) {
	list, err := ImportExcel("test.xlsx")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(list)
	// [map[年龄:男 性别:1 用户名:测试] map[年龄:男 性别:1 用户名:ggr1]]
}
