package xlsx

import (
	"fmt"
	_errors "github.com/pkg/errors"

	"github.com/xuri/excelize/v2"
)

type export struct {
	a1         []string
	rows       [][]interface{}
	a1Length   int
	rowsLength int
	sheetName  string
}

func NewExport() *export {
	return &export{}
}

func (e *export) SetSheetName(name string) *export {
	e.sheetName = name
	return e
}

func (e *export) SetA1(a1 []string) *export {
	e.a1 = a1
	e.a1Length = len(e.a1)
	return e
}

func (e *export) SetRows(rows [][]interface{}) *export {
	e.rows = rows
	e.rowsLength = len(e.rows)
	return e
}

// Excel 导出Excel文件
// headers 列名切片， 表头
// rows 数据切片，是一个二维数组
func (e *export) Excel() (*excelize.File, error) {
	excel := excelize.NewFile()
	excel.SetSheetName("Sheet1", e.sheetName)
	if err := excel.SetSheetRow(e.sheetName, "A1", &e.a1); err != nil {
		return excel, _errors.Wrap(err, "A1行数据 填充失败!")
	}

	for j := 0; j < e.rowsLength; j++ {
		index := fmt.Sprintf("A%d", j+2)
		if err := excel.SetSheetRow(e.sheetName, index, &e.rows[j]); err != nil {
			return excel, _errors.Wrap(err, "excel.SetSheetRow() 填充数据失败!")
		}
	}
	return excel, nil
}
