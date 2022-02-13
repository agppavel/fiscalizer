package fiscalizer

import "time"

type document struct {
	IdDomain    int              `json:"IdDomain"`
	Cash        float32          `json:"Cash"`
	NonCash     float32          `json:"NonCash"`
	Positions   []position       `json:"Positions"`
	Total       float32          `json:"Total"`
	AFP         uint32           `json:"AFP"`
	ReceiptDate time.Time        `json:"ReceiptDate"`
	PngBill     bool             `json:"GenerateCheck"`
	doc_type    fiscal_operation `json:"-"`
}

func NewDocument(operation fiscal_operation, domain int, cash, non_cash, total float32, png_bill bool) document {
	return Document{IdDomain: domain, Cash: cash, NonCash, non_cash, Total: Total, PngBill: png_bill, doc_type: operation}
}

func (self *document) AddPosition(name string, section int, price, marckup, discount, qty float32) {
	self.Positions = append(self.Positions, newPosition(name, section, price, marckup, discount, qty))
}
