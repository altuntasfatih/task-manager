package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(userService interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Create User")
	}
}
func DeleteUser(userService interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("Delete User -> %s", ctx.Params("userId")))
	}
}
func GetUser(userService interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("Get User -> %s", ctx.Params("userId")))
	}
}
func GetUsers(userService interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Get All Users")
	}
}
