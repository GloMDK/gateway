package service

type PayRequest struct {
	Amount       float64
	CurrencyCode uint16
}

type PayResponse struct {
	PayID  string
	Status PayStatus
}

type PayStatusRequest struct {
	PayID string
}

type PayStatusResponse struct {
	Status PayStatus
}

type PayStatus uint8

const (
	PayStatusFail PayStatus = iota
	PayStatusSuccess
	PayStatusPending
)
