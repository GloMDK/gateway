package service

type PayRequest struct {
	Amount       float64 `json:"amount"`
	CurrencyCode uint16  `json:"currency_code"`
	Pan          string  `json:"pan"`
	CVV          string  `json:"CVV"`
	Expired      string  `json:"expired"`
}

type PayResponse struct {
	PayID  string    `json:"pay_id"`
	Status PayStatus `json:"status"`
}

type PayStatusRequest struct {
	PayID string
}

type PayStatus uint8

const (
	PayStatusFail PayStatus = iota
	PayStatusSuccess
	PayStatusPending
	PayStatusNew
)

type TransactionUpdateRequest struct {
	Status PayStatus
	PayID  string
}

type CreateTransactionRequest struct {
	Amount       float64
	CurrencyCode uint16
	BankName     string
}

type GetTransactionResponse struct {
	Status   PayStatus
	BankName string
}
