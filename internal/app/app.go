package app

import (
	"context"
	"github.com/altuntasfatih/task-manager/internal/repository"
	"github.com/altuntasfatih/task-manager/internal/user/handler"
	"github.com/altuntasfatih/task-manager/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"sync"
)

type App struct {
	router      *fiber.App
	userService service.UserService
}

type initializerFunc func() error

func NewApp() (*App, error) {
	app := App{router: fiber.New()}
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
	router := fiber.New()
	router.Use(recover.New())
	router.Use(logger.New())

	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	prefix := "/v1/wallets"
	users := router.Group(prefix)

	users.Post("", handler.CreateUser(""))
	users.Get("", handler.GetUsers(""))
	users.Get("/:userId", handler.GetUser(""))
	users.Delete("/:userId", handler.DeleteUser(""))
	a.router = router
	return nil
}
func (a *App) initService() error {
	repository, err := repository.NewUserRepository()
	if err != nil {
		return err
	}

	userService, err := service.NewUserService(repository)
	if err != nil {
		return err
	}
	a.userService = userService
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
