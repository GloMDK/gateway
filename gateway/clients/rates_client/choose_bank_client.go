package rates_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway/service"
	"io"
	"net/http"
)

const chooseBankClientURLPath = "choose_bank_name"

func (r *RatesClient) ChooseBankClient(ctx context.Context, req *service.ChooseBankClientRequest) (service.BankClient, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.host+"/"+chooseBankClientURLPath, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext error: %w", err)
	}

	resp, err := r.httpClient.Do(httpReq)
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

	bankName := service.BankName(respBody)
	bankClient, found := r.banksClients[bankName]
	if !found {
		return nil, fmt.Errorf("there is no bankClient with name: %s", bankName)
	}

	return bankClient, nil
}
