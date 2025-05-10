package fast_bank

import (
	"fmt"
	"gateway/gateway/service"
	"math/rand/v2"
	"net/http"
)

type SlowBankClient struct {
	httpClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func (c *SlowBankClient) GetBankName() string {
	return "SlowBank"
}

func New(httpClient HTTPClient) *SlowBankClient {
	return &SlowBankClient{httpClient: httpClient}
}

func getRandomStatusOrError() (service.PayStatus, error) {
	randInt := rand.IntN(4)
	if randInt == 3 {
		return 0, fmt.Errorf("random bank error")
	}
	return service.PayStatus(randInt), nil
}
