package fast_bank

import (
	"context"
	"gateway/service"
	"time"
)

func (c *FastBankClient) PayStatus(ctx context.Context, req *service.PayStatusRequest) (service.PayStatus, error) {
	time.Sleep(5 * time.Second)
	return getRandomStatusOrError()
}
