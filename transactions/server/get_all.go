package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetAll(c *fiber.Ctx) error {
	var allTrans []Transactions
	s.db.Find(&allTrans)

	body, err := json.Marshal(allTrans)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("json.Marshal error: %v", err))
	}

	return c.Status(fiber.StatusOK).Send(body)
}
