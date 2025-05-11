package server

import (
	"context"
)

type Server struct {
	cache CacheClient
}

func New(cache CacheClient) *Server {
	return &Server{cache: cache}
}

type CacheClient interface {
	Set(ctx context.Context, key string, val string) error
	Get(ctx context.Context, key string) (string, error)
}
