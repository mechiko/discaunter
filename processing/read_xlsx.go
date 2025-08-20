package processing

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (k *Processing) ReadXlsx(file string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic ReadXlsx %v", r)
		}
	}()

	f, err := excelize.OpenFile(file)
	if err != nil {
		return fmt.Errorf("open xlsx error %w", err)
	}
	defer func() {
		// Close the spreadsheet.
		if errr := f.Close(); errr != nil {
			err = errr
		}
	}()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return fmt.Errorf("len sheet wrong")
	}
	currentSheet := sheets[0]

	// Получить все строки в Sheet1
	rows, err := f.GetRows(currentSheet)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for rowNumber, row := range rows {
		if isRecord(row) {
			// берем запись
			rec, err := NewRecord(row)
			if err != nil {
				return fmt.Errorf("error read xlsx row %d %v", rowNumber+1, err)
			}
			if _, ok := k.Boxes[rec.Box]; !ok {
				k.Boxes[rec.Box] = make([]*Record, 0)
			}
			k.Boxes[rec.Box] = append(k.Boxes[rec.Box], rec)
		}
	}
	return nil
}
