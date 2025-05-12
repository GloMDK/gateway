package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (s *Server) Create(c *fiber.Ctx) error {
	newTrans := &Transaction{}
	body := c.Body()
	err := json.Unmarshal(body, newTrans)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	s.db.Create(newTrans)
	resp := strconv.Itoa(newTrans.ID)

	return c.Status(fiber.StatusOK).Send([]byte(resp))
}
