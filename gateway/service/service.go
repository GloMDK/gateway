package service

import (
	"context"
)

type Service struct {
	ratesClient        RatesClient
	transactionsClient TransactionsClient
}

func New(ratesClient RatesClient, transactionsClient TransactionsClient) *Service {
	return &Service{
		ratesClient:        ratesClient,
		transactionsClient: transactionsClient,
	}
}

type BankClient interface {
	Pay(ctx context.Context, req *PayRequest) (PayStatus, error)
	PayStatus(ctx context.Context, req *PayStatusRequest) (PayStatus, error)
	GetBankName() BankName
}

type RatesClient interface {
	ChooseBankClient(ctx context.Context, req *ChooseBankClientRequest) (BankClient, error)
	GetBankClientByName(bankName BankName) (BankClient, error)
}

type TransactionsClient interface {
	Get(ctx context.Context, payID string) (*GetTransactionResponse, error)
	Create(ctx context.Context, req *CreateTransactionRequest) (payID string, err error)
	Update(ctx context.Context, req *UpdateTransactionRequest) error
}
