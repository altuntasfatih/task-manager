package handler

import (
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// CreateTask godoc
// @Summary Create a task for user
// @Description CreateTask
// @ID CreateTask
// @Tags Tasks
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Param request body models.CreateTaskRequest true "Request"
// @Success 200 {object} models.GetTaskResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId}/tasks [post]
func CreateTask(service service.TaskService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request models.CreateTaskRequest
		if err := ctx.BodyParser(&request); err != nil {
			return err
		}
		if err := validator.ValidateRequest(request); err != nil {
			return err
		}
		userId := ctx.Params("userId", "")
		task, err := service.CreateTask(userId, &request)
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetTaskResponse{
			task,
		})
	}
}

// GetTask godoc
// @Summary Get user's a task by taskId
// @Description GetTask
// @ID GetTask
// @Tags Tasks
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Param taskId path string true "taskID"
// @Success 200 {object} models.GetTaskResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId}/tasks/{taskId} [get]
func GetTask(service service.TaskService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId", "")
		taskId, _ := strconv.Atoi(ctx.Params("taskId", "-1"))
		task, err := service.GetTask(userId, taskId)
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetTaskResponse{
			task,
		})
	}
}

// GetTasks godoc
// @Summary Get user's all task
// @Description GetTasks
// @ID GetTasks
// @Tags Tasks
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Success 200 {object} models.GetTasksResponse
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId}/tasks [get]
func GetTasks(service service.TaskService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId", "")
		tasks, err := service.GetTasks(userId)
		if err != nil {
			return err
		}
		return ctx.JSON(&models.GetTasksResponse{
			tasks,
		})
	}
}

// DeleteTask godoc
// @Summary Delete user's a task by taskId
// @Description DeleteTask
// @ID Delete
// @Tags Tasks
// @Accept json
// @Produce json
// @Param userId path string true "userId"
// @Param taskId path string true "taskID"
// @Success 200
// @Failure 400 {object} custom.ErrorResponse
// @Failure 404 {object} custom.ErrorResponse
// @Failure 500 {object} custom.ErrorResponse
// @Router /v1/users/{userId}/tasks/{taskId} [delete]
func DeleteTask(service service.TaskService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId", "")
		taskId, _ := strconv.Atoi(ctx.Params("taskId", "-1"))
		err := service.DeleteTask(userId, taskId)
		if err != nil {
			return err
		}
		return ctx.Send(nil)
	}
}
