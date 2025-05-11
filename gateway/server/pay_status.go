package server

import (
	"encoding/json"
	"fmt"
	"gateway/service"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) PayStatus(c *fiber.Ctx) error {
	serviceReq := &service.PayStatusRequest{}
	err := json.Unmarshal(c.Body(), serviceReq)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("json.Unmarshal error: %v", err))
	}
	resp, err := s.service.PayStatus(c.UserContext(), serviceReq)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("service.PayStatus error: %v", err))
	}
	body, err := json.Marshal(resp)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("json.Marshal error: %v", err))
	}
	return c.Status(fiber.StatusOK).Send(body)
}
