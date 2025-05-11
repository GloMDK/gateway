package rates_client

import (
	"gateway/service"
	"net/http"
)

type RatesClient struct {
	host         string
	banksClients map[service.BankName]service.BankClient
	httpClient   HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func New(host string, httpClient HTTPClient, banksClients ...service.BankClient) *RatesClient {
	banksClientsMap := map[service.BankName]service.BankClient{}
	for _, client := range banksClients {
		banksClientsMap[client.GetBankName()] = client
	}
	return &RatesClient{
		host:         host,
		httpClient:   httpClient,
		banksClients: banksClientsMap,
	}
}
