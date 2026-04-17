package cli

import (
	"green-api-web-ui/internal/config"
	"green-api-web-ui/internal/container"
	"green-api-web-ui/internal/delivery/http/routes"
	"green-api-web-ui/internal/handler"
	"green-api-web-ui/internal/infrastructure/http"
	"green-api-web-ui/internal/usecase/green_api"

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

	httpClient := http.NewClient(app.container.Config.HTTPClient.BaseURL, app.container.Config.HTTPClient.Timeout)
	useCase := green_api.NewUseCase(httpClient)
	apiHandler := handler.NewAPIHandler(useCase)

	hh := handler.NewHomeHandler()
	routes.Setup(fb, hh, apiHandler)
	app.fiber = fb

	return fb
}
