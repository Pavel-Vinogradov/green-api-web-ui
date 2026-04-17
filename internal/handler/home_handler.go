package handler

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
}
