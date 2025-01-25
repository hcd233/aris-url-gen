package handler

import (
	"github.com/gofiber/fiber/v2"
)

type healthCheckHandler struct{}

// HealthCheck 健康检查
//
// @Summary 健康检查
// @Description 检查服务是否正常运行
// @Tags 系统
// @Accept json
// @Produce plain
// @Success 200 {string} string "ok"
// @Router /health [get]
func (h *healthCheckHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

// NewHealthCheckHandler 创建健康检查接口
//
//	author centonhuang
//	update 2024-12-05 16:09:35
func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}
