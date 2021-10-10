package app

import (
	"context"
	"github.com/altuntasfatih/task-manager/internal/user/handler"
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/altuntasfatih/task-manager/pkg/storage/badger_storage"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"sync"
)

type App struct {
	router      *fiber.App
	userService service.UserService
	taskService service.TaskService
}

type initializerFunc func() error

func NewApp() (*App, error) {
	app := App{router: fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})}
	initializers := []initializerFunc{
		app.initService,
		app.initRouter,
	}
	for _, fn := range initializers {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	return &app, nil
}

func (a *App) initRouter() error {
	router := a.router
	router.Use(recover.New())
	router.Use(logger.New())

	router.Get("/_monitoring/health", handler.HealthCheck())
	router.Get("/swagger/*", fiberSwagger.Handler)
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
	})

	{
		userPrefix := "/v1/users"
		users := router.Group(userPrefix)
		users.Post("", handler.CreateUser(a.userService))
		users.Get("", handler.GetUsers(a.userService))
		users.Get("/:userId", handler.GetUser(a.userService))
		users.Delete("/:userId", handler.DeleteUser(a.userService))
		users.Put("/:userId/reminder", handler.SetReminder(a.userService))
	}

	{
		tasksPrefix := "/v1/users/:userId/tasks"
		tasks := router.Group(tasksPrefix)
		tasks.Post("", handler.CreateTask(a.taskService))
		tasks.Get("", handler.GetTasks(a.taskService))
		tasks.Get("/:taskId", handler.GetTask(a.taskService))
		tasks.Delete("/:taskId", handler.DeleteTask(a.taskService))
	}

	return nil
}

func (a *App) initService() error {
	userStore, err := badger_storage.NewClient(false)
	if err != nil {
		return err
	}

	userService, err := service.NewUserService(userStore)
	if err != nil {
		return err
	}
	taskService, err := service.NewTaskService(userStore)
	if err != nil {
		return err
	}

	a.userService = userService
	a.taskService = taskService
	return nil

}

func (a *App) Listen(ctx context.Context, wg *sync.WaitGroup) {
	log.Println("Web server starting...")

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := "8080"
		if err := a.router.Listen("0.0.0.0:" + port); err != nil {
			log.Fatal(err)
		}

		log.Println("Web server stopped.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		log.Println("Web Server shutting down.")
		a.router.Shutdown()
	}()

}
