package handler

import (
	"github.com/gofiber/fiber/v2"
)

type healthCheckHandler struct{}

// HealthCheck 健康检查
//
//	@author centonhuang
//	@update 2024-12-05 16:09:31
func (h *healthCheckHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

// NewHealthCheckHandler 创建健康检查接口
//
//	@author centonhuang
//	@update 2024-12-05 16:09:35
func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}
