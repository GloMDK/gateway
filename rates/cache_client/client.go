package cache_client

import (
	"context"
	"fmt"
	"sync"
)

type InMemoryCacheClient struct {
	memory map[string][]byte
	mutex  *sync.RWMutex
}

func New() *InMemoryCacheClient {
	return &InMemoryCacheClient{
		memory: map[string][]byte{
			"rates": []byte(`
{
    "rates": {
        "123": [
            {
                "bank_name": "FastBank",
                "rate_value": 0.3
            },
            {
                "bank_name": "SlowBank",
                "rate_value": 0.5
            }
        ],
		"321": [
            {
                "bank_name": "UnknownBank",
                "rate_value": 1
            }
        ]
    }
}
`),
		},
		mutex: &sync.RWMutex{},
	}
}

func (c *InMemoryCacheClient) Get(_ context.Context, key string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	data, found := c.memory[key]
	if !found {
		return nil, fmt.Errorf("no data found by key: %v", key)
	}

	return data, nil
}

func (c *InMemoryCacheClient) Set(_ context.Context, key string, val []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.memory[key] = val

	return nil
}
