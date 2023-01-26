package handler

//
//func initUserRouter(service service.RepairService) *fiber.App {
//	router := fiber.New(fiber.Config{
//		ErrorHandler: ErrorHandler,
//	})
//	prefix := "/v1/repairs"
//	users := router.Group(prefix)
//	users.Post("", CreateRepair(service))
//	users.Get("", GetRepairs(service))
//	users.Get("/:repairId", GetRepair(service))
//	users.Delete("/:repairId", DeleteRepair(service))
//	return router
//}
//
//func TestCreateUser(t *testing.T) {
//	//userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewRepairService(nil)
//	router := initUserRouter(userService)
//
//	request := models.CreateRepairRequest{Email: "test@gmail.com", FirstName: "testName", LastName: "testLastName"}
//	req := createRequest("POST", "/v1/repair", request)
//
//	//when
//	resp, err := router.Test(req)
//
//	//then
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusOK)
//
//	var response models.GetRepairResponse
//	parseResponseToStruct(resp, &response)
//
//	require.Equal(t, response.Email, request.Email)
//	require.Equal(t, response.FirstName, request.FirstName)
//	require.Equal(t, response.LastName, request.LastName)
//	require.NotEmpty(t, response.Id)
//}
//
//func TestCreateUser_WhenRequestInvalid(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//
//	request := models.CreateUserRequest{Email: "test", FirstName: "testName", LastName: "testLastName"}
//	req := createRequest("POST", "/v1/users", request)
//
//	//when
//	resp, err := router.Test(req)
//
//	//then
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusBadRequest)
//}
//
//func TestGetUser(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//
//	expectedUser := models.NewUser("1", "test@gmail.com", "testName", "testLastName")
//	storeUser(userStore, expectedUser)
//	req := createRequest("GET", "/v1/users/"+expectedUser.Id, nil)
//
//	//when
//	resp, err := router.Test(req)
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusOK)
//
//	var response models.GetUserResponse
//	parseResponseToStruct(resp, &response)
//
//	require.Equal(t, expectedUser.Email, response.Email)
//	require.Equal(t, expectedUser.FirstName, response.FirstName)
//	require.Equal(t, expectedUser.LastName, response.LastName)
//	require.NotEmpty(t, expectedUser.Id)
//}
//
//func TestGetUser_WhenUserNotFound(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//
//	req := createRequest("GET", "/v1/users/fake", nil)
//
//	//when
//	resp, err := router.Test(req)
//
//	//then
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusNotFound)
//}
//
//func TestGetUsers(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//
//	storeUser(userStore, models.NewUser("1", "test@gmail.com", "testName", "testLastName"))
//	storeUser(userStore, models.NewUser("2", "test@gmail.com", "testName", "testLastName"))
//	storeUser(userStore, models.NewUser("3", "test@gmail.com", "testName", "testLastName"))
//
//	req := createRequest("GET", "/v1/users/", nil)
//
//	//when
//	resp, err := router.Test(req)
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusOK)
//
//	var response models.GetUsersResponse
//	parseResponseToStruct(resp, &response)
//
//	require.Equal(t, len(response.Users), 3)
//}
//
//func TestDeleteUser(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//	storeUser(userStore, models.NewUser("1", "test@gmail.com", "testName", "testLastName"))
//
//	req := createRequest("DELETE", "/v1/users/1", nil)
//
//	//when
//	resp, err := router.Test(req)
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusOK)
//}
//
//func TestSetReminder(t *testing.T) {
//	userStore, _ := badger_storage.NewClient(true)
//	userService, _ := service.NewUserService(userStore)
//	router := initUserRouter(userService)
//	storeUser(userStore, models.NewUser("1", "test@gmail.com", "testName", "testLastName"))
//
//	request := models.SetReminderRequest{Method: models.Onsite}
//	req := createRequest("PUT", "/v1/users/1/reminder", request)
//
//	//when
//	resp, err := router.Test(req)
//
//	//then
//	require.Nil(t, err)
//	require.Equal(t, resp.StatusCode, fiber.StatusOK)
//
//	var response models.GetUserResponse
//	parseResponseToStruct(resp, &response)
//
//	require.Equal(t, response.ReminderMethod, models.Onsite)
//
//}
//
//func storeUser(store storage.Writer, user *models.User) {
//	_ = store.CreateUser(user.Id, user)
//}
//
//func convertRequest(request interface{}) io.Reader {
//	requestBody, _ := json.Marshal(request)
//	return bytes.NewReader(requestBody)
//}
//
//func parseResponseToStruct(response *http.Response, value interface{}) {
//	body, _ := ioutil.ReadAll(response.Body)
//	_ = json.Unmarshal(body, value)
//}
//
//func createRequest(method, router string, request interface{}) *http.Request {
//	req := httptest.NewRequest(method, router, convertRequest(request))
//	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
//	return req
//}
