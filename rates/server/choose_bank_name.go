package server

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) ChooseBankName(c *fiber.Ctx) error {
	req := &ChooseBankNameRequest{}
	err := json.Unmarshal(c.Body(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	paramBody, err := s.cache.Get(c.UserContext(), RatesParameterCacheKey)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("cache.Set error: %v", err))
	}
	param := &RatesParameter{}
	err = json.Unmarshal(paramBody, param)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("json.Unmarshal error: %v", err))
	}

	banks, found := param.Rates[req.CurrencyCode]
	if !found {
		return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("no such currency in parameter: %v", req.CurrencyCode))
	}
	if len(banks) == 0 {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("there are no banks for chosen currency: %v", req.CurrencyCode))
	}

	minRate := banks[0].RateValue
	bankName := banks[0].BankName
	for _, bank := range banks {
		if bank.RateValue < minRate {
			minRate = bank.RateValue
			bankName = bank.BankName
		}
	}

	return c.Status(fiber.StatusOK).Send([]byte(bankName))
}
