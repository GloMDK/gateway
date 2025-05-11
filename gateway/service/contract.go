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

type UpdateTransactionRequest struct {
	Status PayStatus `json:"status"`
	PayID  string    `json:"-"`
}

type CreateTransactionRequest struct {
	Amount       float64   `json:"amount"`
	CurrencyCode uint16    `json:"currency_code"`
	BankName     BankName  `json:"bank_name"`
	Status       PayStatus `json:"status"`
}

type GetTransactionResponse struct {
	Status   PayStatus `json:"status"`
	BankName BankName  `json:"bank_name"`
}

type BankName string

type ChooseBankClientRequest struct {
	Pan          string  `json:"pan"`
	Amount       float64 `json:"amount"`
	CurrencyCode uint16  `json:"currency_code"`
}
