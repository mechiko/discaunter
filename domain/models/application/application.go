package application

type Zakaz struct {
	Magazin  string
	Quantity string
	Gtin     string
	ID       string
	CodeAP   string // ищем в А3 в справочнике для контроля
}
