package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Update(c *fiber.Ctx) error {
	req := &UpdateRequest{}
	body := c.Body()
	err := json.Unmarshal(body, req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	s.db.Model(&Transactions{}).Where("id = ?", c.Params("+")).Update("status", req.Status)

	return c.SendStatus(fiber.StatusOK)
}
