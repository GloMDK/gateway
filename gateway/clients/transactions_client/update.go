package transactions_client

import (
	"context"
	"fmt"
	"gateway/service"
	"net/http"
)

func (t *TransactionsClient) Update(ctx context.Context, req *service.UpdateTransactionRequest) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPatch, t.host+"/"+req.PayID, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext error: %w", err)
	}

	resp, err := t.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("httpClient.Do error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("httpClient.Do status code is: %v", resp.StatusCode)
	}

	return nil
}
