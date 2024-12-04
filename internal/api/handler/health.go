package handler

import (
	"github.com/gofiber/fiber/v2"
)

type healthCheckHandler struct{}

func (h *healthCheckHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}
