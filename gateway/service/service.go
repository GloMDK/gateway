package service

type Service struct {
	ratesClient        RatesClient
	transactionsClient TransactionsClient
}

type BankClient interface {
	Pay() error
	PayStatus() error
}

type RatesClient interface {
	ChooseBankClient() (BankClient, error)
	GetBankClient() (BankClient, error)
}

type TransactionsClient interface {
	Create() error
	Update() error
}
