package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"status": "UP",
		})
	}
}
