package fast_bank

import (
	"context"
	"gateway/gateway/service"
	"time"
)

func (c *SlowBankClient) Pay(ctx context.Context, req *service.PayRequest) (service.PayStatus, error) {
	time.Sleep(30 * time.Second)
	return getRandomStatusOrError()
}
