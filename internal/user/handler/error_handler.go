package handler

import (
	"github.com/altuntasfatih/task-manager/pkg/custom"
	"github.com/altuntasfatih/task-manager/pkg/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	if err, ok := err.(validator.ValidationErrors); ok {
		return c.Status(fiber.StatusBadRequest).JSON(custom.ErrorResponse{Message: err.Error()})

	}
	if err == store.ErrUserNotFound {
		return c.Status(fiber.StatusNotFound).JSON(custom.ErrorResponse{Message: err.Error()})

	}
	return c.Status(fiber.StatusInternalServerError).JSON(custom.ErrorResponse{Message: err.Error()})
}