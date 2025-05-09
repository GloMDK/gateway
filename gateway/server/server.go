package server

import "gateway/gateway/service"

type Server struct {
	service Service
}

func New(service Service) *Server {
	return &Server{service: service}
}

type Service interface {
	Pay(req *service.PayRequest) (*service.PayResponse, error)
	PayStatus(req *service.PayStatusRequest) (*service.PayStatusResponse, error)
}
