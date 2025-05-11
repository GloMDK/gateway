package slow_bank

import (
	"context"
	"gateway/service"
	"time"
)

func (c *SlowBankClient) Pay(ctx context.Context, req *service.PayRequest) (service.PayStatus, error) {
	time.Sleep(29 * time.Second)
	return getRandomStatusOrError()
}
