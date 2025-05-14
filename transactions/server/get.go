package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) Get(c *fiber.Ctx) error {
	trans := &Transactions{}
	s.db.First(trans, "id = ?", c.Params("+"))
	if trans.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	body, err := json.Marshal(trans)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("json.Marshal error: %v", err))
	}

	return c.Status(fiber.StatusOK).Send(body)
}
