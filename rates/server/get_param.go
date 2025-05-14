package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetParam(c *fiber.Ctx) error {
	paramBody, err := s.cache.Get(c.UserContext(), RatesParameterCacheKey)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("cache.Set error: %v", err))
	}

	newParam := &RatesParameter{}
	err = json.Unmarshal(paramBody, newParam)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	return c.Status(fiber.StatusOK).Send(paramBody)
}
