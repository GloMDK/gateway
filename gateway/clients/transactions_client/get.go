package transactions_client

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/service"
	"io"
	"net/http"
)

func (t *TransactionsClient) Get(ctx context.Context, payID string) (*service.GetTransactionResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, t.host+"/"+payID, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext error: %w", err)
	}

	resp, err := t.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("httpClient.Do status code is: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	finalResp := &service.GetTransactionResponse{}
	err = json.Unmarshal(respBody, finalResp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return finalResp, nil
}
