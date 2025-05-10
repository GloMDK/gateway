package fast_bank

import (
	"context"
	"gateway/gateway/service"
	"time"
)

func (c *FastBankClient) Pay(ctx context.Context, req *service.PayRequest) (service.PayStatus, error) {
	time.Sleep(5 * time.Second)
	return getRandomStatusOrError()
}
