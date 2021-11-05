package snowflake

import (
	"fmt"
	"testing"
)

func Test_Id(t *testing.T) {
	str := Id()
	fmt.Println(str)
}