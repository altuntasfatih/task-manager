package handler

import (
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Create User
// @Description CreateUser
// @ID UserCreate
// @Tags Users
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Request"
// @Success 200 {object} models.GetUserResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users [post]
func CreateUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request models.CreateUserRequest
		if err := ctx.BodyParser(&request); err != nil {
			return err
		}

		if err := validator.ValidateRequest(request); err != nil {
			return err
		}

		user, err := service.CreateUser(&request)
		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
}

// GetUser godoc
// @Summary Get User by userId
// @Description GetUser
// @ID GetUser
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Success 200 {object} models.GetUserResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId} [get]
func GetUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId", "")
		user, err := service.GetUser(userId)
		if err != nil {
			return err
		}

		return ctx.JSON(&models.GetUserResponse{
			User: user,
		})
	}
}

// GetUsers godoc
// @Summary Get All Users
// @Description GetUsers
// @ID GetUsers
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Success 200 {object} models.GetUsersResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users [get]
func GetUsers(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := service.GetUsers()
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetUsersResponse{
			Users: users,
		})
	}
}

// DeleteUser godoc
// @Summary Delete user by userId
// @Description DeleteUser
// @ID DeleteUser
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Success 200
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId} [delete]
func DeleteUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId", "")
		if err := service.DeleteUser(userId); err != nil {
			return err
		}
		return ctx.Send(nil)
	}
}
