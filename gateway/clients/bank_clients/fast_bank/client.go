package fast_bank

import (
	"fmt"
	"gateway/gateway/service"
	"math/rand/v2"
	"net/http"
)

type FastBankClient struct {
	httpClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func (c *FastBankClient) GetBankName() string {
	return "FastBank"
}

func New(httpClient HTTPClient) *FastBankClient {
	return &FastBankClient{httpClient: httpClient}
}

func getRandomStatusOrError() (service.PayStatus, error) {
	randInt := rand.IntN(4)
	if randInt == 3 {
		return 0, fmt.Errorf("random bank error")
	}
	return service.PayStatus(randInt), nil
}
