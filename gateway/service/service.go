package service

import (
	"context"
)

type Service struct {
	ratesClient        RatesClient
	transactionsClient TransactionsClient
}

type BankClient interface {
	Pay(ctx context.Context, req *PayRequest) (PayStatus, error)
	PayStatus(ctx context.Context, req *PayStatusRequest) (PayStatus, error)
	GetBankName() string
}

type RatesClient interface {
	ChooseBankClient(ctx context.Context, req *PayRequest) (BankClient, error)
	GetBankClientByName(ctx context.Context, bankName string) (BankClient, error)
}

type TransactionsClient interface {
	GetTransaction(ctx context.Context, payID string) (*GetTransactionResponse, error)
	Create(ctx context.Context, req *CreateTransactionRequest) (payID string, err error)
	Update(ctx context.Context, req *TransactionUpdateRequest) error
}
