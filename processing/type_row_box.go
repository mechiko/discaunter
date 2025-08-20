package processing

import "fmt"

type Record struct {
	Mark string
	Box  string
}

func NewRecord(row []string) (*Record, error) {
	if len(row) < 11 {
		return nil, fmt.Errorf("row len %d less then 11", len(row))
	}
	r := &Record{
		Mark: row[9],
		Box:  row[10],
	}
	return r, nil
}

// полная строка 11 ячеек
func isRecord(row []string) bool {
	return len(row) > 10 && row[1] != "Ссылка"
}
