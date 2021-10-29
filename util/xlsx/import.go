package xlsx

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

func ImportExcel(fullPath string) ([]map[string]string, error) {
	// 处理上传
	f, err := excelize.OpenFile(fullPath)
	if err != nil {
		return nil, err
	}
	// 获取 Sheet1 上所有单元格
	importData, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	if len(importData) == 0 {
		return nil, errors.New("xlsx is empty")
	}

	title := importData[0]

	var res []map[string]string

	for _, row := range importData[1:] {
		var infoMap = make(map[string]string)

		for t := 0; t < len(title); t++ {
			infoMap[title[t]] = row[t]
		}

		res = append(res, infoMap)
	}

	return res, nil
}
