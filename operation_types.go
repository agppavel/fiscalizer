package fiscalizer

type fiscal_operation struct {
	endpoint string
	method   string
}

const (
	SALE            fiscal_operation = fiscal_operation{endpoint: "kkms/%d/sales", method: "POST"}
	PURCHASE        fiscal_operation = fiscal_operation{endpoint: "kkms/%d/purchases", method: "POST"}
	REFUND          fiscal_operation = fiscal_operation{endpoint: "kkms/%d/refunds", method: "POST"}
	PURCHASE_REFUND fiscal_operation = fiscal_operation{endpoint: "kkms/%d/purchase_refunds", method: "POST"}
)
