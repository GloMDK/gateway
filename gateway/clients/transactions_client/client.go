package transactions_client

import "net/http"

type TransactionsClient struct {
	host       string
	httpClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func New(host string, httpClient HTTPClient) *TransactionsClient {
	return &TransactionsClient{
		host:       host,
		httpClient: httpClient,
	}
}
