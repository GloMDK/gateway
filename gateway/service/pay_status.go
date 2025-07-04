package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Service) PayStatus(ctx context.Context, req *PayStatusRequest) (PayStatus, error) {
	transaction, err := s.transactionsClient.Get(ctx, req.PayID)
	if err != nil {
		return 0, fmt.Errorf("payID: %v, transactionsClient.Get error: %w", req.PayID, err)
	}
	if transaction.Status == PayStatusSuccess || transaction.Status == PayStatusFail {
		return transaction.Status, nil
	}

	bankClient, err := s.ratesClient.GetBankClientByName(transaction.BankName)
	if err != nil {
		return transaction.Status, fmt.Errorf("ratesClient.GetBankClientByName error: %w", err)
	}
	bankStatus, err := bankClient.PayStatus(ctx, req)
	if err != nil {
		log.Info(fmt.Sprintf("payID: %v, bankClient.PayStatus error: %v", req.PayID, err))
		return transaction.Status, nil
	}

	if bankStatus != transaction.Status {
		err = s.transactionsClient.Update(ctx, &UpdateTransactionRequest{
			Status: bankStatus,
			PayID:  req.PayID,
		})
		if err != nil {
			log.Info(fmt.Sprintf("payID: %v, transactionsClient.Update error: %v", req.PayID, err))
		}
	}

	return bankStatus, nil
}
