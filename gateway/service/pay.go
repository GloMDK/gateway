package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Service) Pay(ctx context.Context, req *PayRequest) (*PayResponse, error) {
	bankClient, err := s.ratesClient.ChooseBankClient(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ratesClient.ChooseBankClient error: %w", err)
	}

	payID, err := s.transactionsClient.Create(ctx, &CreateTransactionRequest{
		Amount:       req.Amount,
		CurrencyCode: req.CurrencyCode,
		BankName:     bankClient.GetBankName(),
	})
	if err != nil {
		return nil, fmt.Errorf("transactionsClient.Create error: %w", err)
	}

	payStatus, err := bankClient.Pay(ctx, req)
	if err != nil {
		updateErr := s.transactionsClient.Update(ctx, &TransactionUpdateRequest{
			Status: PayStatusFail,
			PayID:  payID,
		})
		if updateErr != nil {
			log.Info(fmt.Sprintf("payID: %v, transactionsClient.Update error: %v", payID, err))
		}
		return nil, fmt.Errorf("bankClient.Pay error: %w", err)
	}

	err = s.transactionsClient.Update(ctx, &TransactionUpdateRequest{
		Status: payStatus,
		PayID:  payID,
	})
	if err != nil {
		log.Info(fmt.Sprintf("payID: %v, transactionsClient.Update error: %v", payID, err))
	}

	return &PayResponse{
		PayID:  payID,
		Status: payStatus,
	}, nil
}
