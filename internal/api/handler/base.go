package handler

import "github.com/gofiber/fiber/v2"

// HealthCheckHandler 健康检查接口
//
//	@author centonhuang
//	@update 2024-12-05 16:09:19
type HealthCheckHandler interface {
	HealthCheck(ctx *fiber.Ctx) error
}

// ShortURLHandler 短URL接口
//
//	@author centonhuang
//	@update 2024-12-05 16:09:23
type ShortURLHandler interface {
	GenerateShortURL(ctx *fiber.Ctx) error
	GetOriginalURL(ctx *fiber.Ctx) error
}
