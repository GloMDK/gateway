package server

const RatesParameterCacheKey = "rates"

type RatesParameter struct {
	Rates map[CurrencyCodeString][]Rates `json:"rates"`
}

type Rates struct {
	BankName  BankName  `json:"bank_name"`
	RateValue RateValue `json:"rate_value"`
}

type ChooseBankNameRequest struct {
	CurrencyCode CurrencyCodeUint `json:"currency_code"`
}

type RateValue float64
type BankName string
type CurrencyCodeUint uint16
type CurrencyCodeString string
