package xlsx

import (
	"fmt"
	"testing"
)

func TestExportExcel(t *testing.T) {

	t.Run("", func(t *testing.T) {
		a1 := []string{"用户名", "性别", "年龄"}
		rows := [][]interface{}{
			{"测试", "1", "男"},
			{"ggr1", "1", "男"},
		}
		_export := NewExport()
		file, err := _export.SetA1(a1).SetSheetName("Sheet1").SetRows(rows).Excel()
		if err != nil {
			t.Error(err)
			return
		}
		if err = file.SaveAs("test.xlsx"); err != nil {
			t.Error(err)
			return
		}
		fmt.Sprintln(err)
	})

}
