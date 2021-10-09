package handler

import (
	"bytes"
	"encoding/json"
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/altuntasfatih/task-manager/pkg/store"
	"github.com/altuntasfatih/task-manager/pkg/store/badger_store"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initRouter(service service.UserService) *fiber.App {
	router := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	prefix := "/v1/users"
	users := router.Group(prefix)
	users.Post("", CreateUser(service))
	users.Get("", GetUsers(service))
	users.Get("/:userId", GetUser(service))
	users.Delete("/:userId", DeleteUser(service))
	return router
}

func TestCreateUser(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	request := models.CreateUserRequest{Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"}
	req := createRequest("POST", "/v1/users", request)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var user models.User
	parseResponseToStruct(resp, &user)

	require.Equal(t, user.Email, request.Email)
	require.Equal(t, user.FirstName, request.FirstName)
	require.Equal(t, user.LastName, request.LastName)
	require.NotEmpty(t, user.Id)
}

func TestCreateUser_WhenRequestInvalid(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	request := models.CreateUserRequest{Email: "test", FirstName: "testName", LastName: "testLastName"}
	req := createRequest("POST", "/v1/users", request)

	//when
	resp, err := router.Test(req)

	//then
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusBadRequest)
}

func TestGetUser(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	expectedUser := &models.User{Id: "1", Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"}
	storeUser(userStore, expectedUser)
	req := createRequest("GET", "/v1/users/"+expectedUser.Id, nil)

	//when
	resp, err := router.Test(req)
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var response models.GetUserResponse
	parseResponseToStruct(resp, &response)

	require.Equal(t, expectedUser.Email, response.Email)
	require.Equal(t, expectedUser.FirstName, response.FirstName)
	require.Equal(t, expectedUser.LastName, response.LastName)
	require.NotEmpty(t, expectedUser.Id)
}

func TestGetUser_WhenUserNotFound(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	req := createRequest("GET", "/v1/users/fake", nil)

	//when
	resp, err := router.Test(req)
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusNotFound)
}

func TestGetUsers(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	storeUser(userStore, &models.User{Id: "1", Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"})
	storeUser(userStore, &models.User{Id: "2", Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"})
	storeUser(userStore, &models.User{Id: "3", Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"})

	req := createRequest("GET", "/v1/users/", nil)

	//when
	resp, err := router.Test(req)
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)

	var response models.GetUsersResponse
	parseResponseToStruct(resp, &response)

	require.Equal(t, len(response.Users), 3)
}

func TestDeleteUser(t *testing.T) {
	userStore, _ := badger_store.NewClient(true)
	userService, _ := service.NewUserService(userStore)
	router := initRouter(userService)

	storeUser(userStore, &models.User{Id: "1", Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"})

	req := createRequest("DELETE", "/v1/users/1", nil)

	//when
	resp, err := router.Test(req)
	require.Nil(t, err)
	require.Equal(t, resp.StatusCode, fiber.StatusOK)
}

func storeUser(store store.Writer, user *models.User) {
	_ = store.CreateUser(user.Id, user)
}

func convertRequest(request interface{}) io.Reader {
	requestBody, _ := json.Marshal(request)
	return bytes.NewReader(requestBody)
}

func parseResponseToStruct(response *http.Response, value interface{}) {
	body, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(body, value)
}

func createRequest(method, router string, request interface{}) *http.Request {
	req := httptest.NewRequest(method, router, convertRequest(request))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return req
}
