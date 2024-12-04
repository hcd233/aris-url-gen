package handler

import "github.com/gofiber/fiber/v2"

type HealthCheckHandler interface {
	HealthCheck(ctx *fiber.Ctx) error
}

type ShortURLHandler interface {
	GenerateShortURL(ctx *fiber.Ctx) error
	GetOriginalURL(ctx *fiber.Ctx) error
}
