package cli

import (
	"green-api-web-ui/internal/config"
	"green-api-web-ui/internal/container"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	container *container.Container
	fiber     *fiber.App
}

func NewApp(cfg *config.AppConfig) (*App, error) {
	app := &App{
		container: container.NewContainer(cfg),
	}
	return app, nil
}

func (app *App) RunApi() *fiber.App {

	fb := fiber.New(fiber.Config{
		AppName:      "Green API",
		ServerHeader: "Green API",
	})

	fb.Use(logger.New())
	fb.Use(cors.New())
	fb.Use(recover.New())

	fb.Static("/css", "./public/css")
	fb.Static("/js", "./public/js")

	app.fiber = fb

	return fb
}
