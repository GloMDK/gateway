package server

const RatesParameterCacheKey = "rates"

type RatesParameter struct {
	Rates map[CurrencyCode][]Rates `json:"rates"`
}

type Rates struct {
	BankName  BankName  `json:"bank_name"`
	RateValue RateValue `json:"rate_value"`
}

type RateValue float64
type BankName string
type CurrencyCode uint16
