package server

import (
	"context"
	"gateway/gateway/service"
)

type Server struct {
	service Service
}

func New(service Service) *Server {
	return &Server{service: service}
}

type Service interface {
	Pay(ctx context.Context, req *service.PayRequest) (*service.PayResponse, error)
	PayStatus(ctx context.Context, req *service.PayStatusRequest) (service.PayStatus, error)
}
