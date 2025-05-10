package fast_bank

import (
	"context"
	"gateway/gateway/service"
	"time"
)

func (c *SlowBankClient) PayStatus(ctx context.Context, req *service.PayStatusRequest) (service.PayStatus, error) {
	time.Sleep(30 * time.Second)
	return getRandomStatusOrError()
}
