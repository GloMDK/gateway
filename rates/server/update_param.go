package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) UpdateParam(c *fiber.Ctx) error {
	newParam := &RatesParameter{}
	body := c.Body()
	err := json.Unmarshal(body, newParam)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	err = s.cache.Set(c.UserContext(), RatesParameterCacheKey, string(body))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("cache.Set error: %v", err))
	}

	return c.SendStatus(fiber.StatusOK)
}
