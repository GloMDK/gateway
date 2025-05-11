package slow_bank

import (
	"fmt"
	"gateway/service"
	"math/rand/v2"
	"net/http"
)

const bankName = "SlowBank"

type SlowBankClient struct {
	httpClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func (c *SlowBankClient) GetBankName() service.BankName {
	return bankName
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
