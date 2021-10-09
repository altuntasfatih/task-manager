package handler

import (
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/store"
	"github.com/altuntasfatih/task-manager/pkg/store/badger_store"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

var testUser = models.NewUser("1", "test@gmail.com", "testName", "testLastName")

func initTaskRouter(service service.TaskService) *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	tasksPrefix := "/v1/users/:userId/tasks"
	tasks := router.Group(tasksPrefix)
	tasks.Post("", CreateTask(service))
	tasks.Get("", GetTasks(service))
	tasks.Get("/:taskId", GetTask(service))
	tasks.Delete("/:taskId", DeleteTask(service))
	return router
}

func TestCreateTask(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)

	startTime := time.Now().Add(48 * time.Hour)
	request := &models.CreateTaskRequest{Name: "daily", StartTime: startTime, EndTime: startTime.Add(30 * time.Minute), ReminderPeriodInHour: 10}
	req := createRequest("POST", "/v1/users/"+testUser.Id+"/tasks", request)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var response models.GetTaskResponse
	parseResponseToStruct(resp, &response)

	require.Equal(t, response.Name, request.Name)
	require.True(t, response.StartTime.Equal(request.StartTime))
	require.True(t, response.EndTime.Equal(request.EndTime))
	require.Equal(t, response.ReminderPeriodInHour, request.ReminderPeriodInHour)
	require.NotEmpty(t, response.Id)
}
func TestCreateTask_WhenInvalidRequest(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)

	request := &models.CreateTaskRequest{Name: "daily", StartTime: time.Now(), EndTime: time.Now(), ReminderPeriodInHour: -5}
	req := createRequest("POST", "/v1/users/"+testUser.Id+"/tasks", request)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusBadRequest)
}

func TestGetTask(t *testing.T) {
	startTime := time.Now().Add(48 * time.Hour)
	task := models.NewTask(1, "daily", startTime, startTime.Add(30*time.Minute), 10)
	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)
	storeTask(userStore, testUser.Id, task)

	req := createRequest("GET", "/v1/users/"+testUser.Id+"/tasks/"+strconv.Itoa(task.Id), nil)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var response models.GetTaskResponse
	parseResponseToStruct(resp, &response)

	require.Equal(t, response.Name, task.Name)
	require.True(t, response.StartTime.Equal(task.StartTime))
	require.True(t, response.EndTime.Equal(task.EndTime))
	require.Equal(t, response.ReminderPeriodInHour, task.ReminderPeriodInHour)
	require.Equal(t, response.Id, task.Id)

}

func TestGetTask_WhenTaskNotFound(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)

	req := createRequest("GET", "/v1/users/"+testUser.Id+"/tasks/"+"fake", nil)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusNotFound)
}

func TestGetTasks(t *testing.T) {

	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)
	storeTask(userStore, testUser.Id, models.NewTask(1, "daily", time.Now().Add(48*time.Hour), time.Now().Add(49*time.Hour), 5))
	storeTask(userStore, testUser.Id, models.NewTask(2, "grooming", time.Now().Add(48*time.Hour), time.Now().Add(49*time.Hour), 3))
	storeTask(userStore, testUser.Id, models.NewTask(3, "planning", time.Now().Add(48*time.Hour), time.Now().Add(49*time.Hour), 2))
	storeTask(userStore, testUser.Id, models.NewTask(4, "fake", time.Now().Add(48*time.Hour), time.Now().Add(49*time.Hour), 1))
	req := createRequest("GET", "/v1/users/"+testUser.Id+"/tasks", nil)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var response models.GetTasksResponse
	parseResponseToStruct(resp, &response)

	require.Equal(t, len(response.Tasks), 4)
}

func TestDeleteTask(t *testing.T) {

	userStore, _ := badger_store.NewClient(true)
	taskService, _ := service.NewTaskService(userStore)
	router := initTaskRouter(taskService)
	storeUser(userStore, testUser)
	storeTask(userStore, testUser.Id, models.NewTask(4, "fake", time.Now().Add(48*time.Hour), time.Now().Add(49*time.Hour), 1))
	req := createRequest("DELETE", "/v1/users/"+testUser.Id+"/tasks/"+"4", nil)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	tasks, err := taskService.GetTasks(testUser.Id)

	require.Equal(t, len(tasks), 0)
}

func storeTask(store store.ReaderWriterRemover, userId string, task *models.Task) {
	user, _ := store.GetUser(userId)
	user.Tasks = append(user.Tasks, task)
	_ = store.UpdateUser(user.Id, user)
}
