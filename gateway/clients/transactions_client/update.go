package transactions_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway/service"
	"net/http"
)

func (t *TransactionsClient) Update(ctx context.Context, req *service.UpdateTransactionRequest) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPatch, t.host+"/"+req.PayID, bytes.NewBuffer(reqBody))
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
