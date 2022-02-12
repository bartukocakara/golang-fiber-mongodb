package config

import (
	"go-delivery-food/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config{
	return fiber.Config{
		ErrorHandler : exception.ErrorHandler,
	}
}