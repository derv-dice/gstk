package io

import (
	"github.com/xuri/excelize/v2"
	"sync"
)

const xlsxDefaultSheetName = "Sheet1"

type WriterXLSX struct {
	f         *excelize.File
	rowsCount map[string]int
	mu        sync.Mutex
}

func NewWriterXLSX() *WriterXLSX {
	return &WriterXLSX{
		f: excelize.NewFile(),
		rowsCount: map[string]int{
			xlsxDefaultSheetName: 0,
		},
	}
}

func (w *WriterXLSX) Close() error {
	return w.f.Close()
}

// WriteRow - Добавить строку с произвольными значениями в конец страницы
// Удобнее всего использовать RowXLSX для формирования строки из набора значений
func (w *WriterXLSX) WriteRow(sheetName string, rows ...[]any) (err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if sheetName == "" {
		sheetName = xlsxDefaultSheetName
	}

	var sheetIndex int
	sheetIndex, err = w.f.GetSheetIndex(sheetName)
	if err != nil {
		return
	}

	if sheetIndex == -1 {
		if sheetIndex, err = w.f.NewSheet(sheetName); err != nil {
			return
		}
	}

	w.f.SetActiveSheet(sheetIndex)

	index := w.rowsCount[sheetName]

	for i1 := range rows {
		for i2 := range rows[i1] {
			var cellName string
			if cellName, err = excelize.CoordinatesToCellName(i2+1, index+1); err != nil {
				return
			}

			if err = w.f.SetCellValue(sheetName, cellName, rows[i1][i2]); err != nil {
				return
			}
		}

		index += 1
	}

	w.rowsCount[sheetName] = index

	return
}

func (w *WriterXLSX) SaveAs(filename string) (err error) {
	return w.f.SaveAs(filename)
}

func RowXLSX(v ...any) (res []any) {
	for i := range v {
		res = append(res, v[i])
	}
	return
}
