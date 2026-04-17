package handler

import (
	"fmt"
	"green-api-web-ui/internal/interfaces/green_api"

	"github.com/gofiber/fiber/v2"
)

type APIHandler struct {
	useCase green_api.UseCaseGreenApi
}

func NewAPIHandler(useCase green_api.UseCaseGreenApi) *APIHandler {
	return &APIHandler{
		useCase: useCase,
	}
}

func (h *APIHandler) GetSettings(c *fiber.Ctx) error {
	var req green_api.APIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
	}

	resp, err := h.useCase.GetSettings(req.IDInstance, req.APITokenInstance)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}

func (h *APIHandler) GetStateInstance(c *fiber.Ctx) error {
	var req green_api.APIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
	}

	resp, err := h.useCase.GetStateInstance(req.IDInstance, req.APITokenInstance)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}

func (h *APIHandler) SendMessage(c *fiber.Ctx) error {
	var req green_api.APIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
	}

	resp, err := h.useCase.SendMessage(req.IDInstance, req.APITokenInstance, req.ChatID, req.Message)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}

func (h *APIHandler) SendFileByUrl(c *fiber.Ctx) error {
	var req green_api.APIRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
	}

	resp, err := h.useCase.SendFileByUrl(req.IDInstance, req.APITokenInstance, req.ChatID, req.URLFile, req.FileName, req.Caption)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}
