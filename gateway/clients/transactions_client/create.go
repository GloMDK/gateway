package transactions_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway/service"
	"io"
	"net/http"
)

func (t *TransactionsClient) Create(ctx context.Context, req *service.CreateTransactionRequest) (string, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("json.Marshal error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, t.host, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("http.NewRequestWithContext error: %w", err)
	}

	resp, err := t.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("httpClient.Do error: %w", err)
	}
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return "", fmt.Errorf("httpClient.Do status code is: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.ReadAll error: %w", err)
	}

	return string(respBody), nil
}
