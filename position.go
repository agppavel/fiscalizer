package fiscalizer

type position struct {
	Name      string  `json:"Name" `
	IdSection int     `json:"IdSection"`
	Price     float32 `json:"Price" `
	Markup    float32 `json:"Markup" `
	Discount  float32 `json:"Discount"`
	Qty       float32 `json:"Qty"`
	Storno    bool    `json:"Storno"`
}

func newPosition(name string, section int, price, marckup, discount, qty float32) position {
	return position{Name: name, IdSection: section, Price: price, Markup: marckup, Discount: discount, Qty: qty}

}
