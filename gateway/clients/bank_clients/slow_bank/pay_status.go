package slow_bank

import (
	"context"
	"gateway/service"
	"time"
)

func (c *SlowBankClient) PayStatus(ctx context.Context, req *service.PayStatusRequest) (service.PayStatus, error) {
	time.Sleep(29 * time.Second)
	return getRandomStatusOrError()
}
