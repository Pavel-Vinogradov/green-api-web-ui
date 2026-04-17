package routes

import (
	"green-api-web-ui/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, home *handler.HomeHandler, api *handler.APIHandler) {
	app.Get("/", home.Index)

	apiGroup := app.Group("/api")
	apiGroup.Post("/getSettings", api.GetSettings)
	apiGroup.Post("/getStateInstance", api.GetStateInstance)
	apiGroup.Post("/sendMessage", api.SendMessage)
	apiGroup.Post("/sendFileByUrl", api.SendFileByUrl)
}
