// Package auth 认证中间件
//
//	@update 2024-12-06 20:00:42
package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/dto"
	"github.com/hcd233/Aris-url-gen/internal/config"
)

// New 认证中间件
//
//	@return fiber.Handler
//	@author centonhuang
//	@update 2024-12-06 20:00:42
func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authKey := c.Get("Authorization")
		authKey = strings.TrimPrefix(authKey, "Bearer ")

		if authKey == "" {
			code := dto.CodeUnauthorized
			return c.Status(fiber.StatusUnauthorized).JSON(dto.StandardResponse{
				Code:    code,
				Message: code.ToMessage("Need Authorization"),
			})
		}
		if authKey != config.AuthKey {
			code := dto.CodeForbidden
			return c.Status(fiber.StatusForbidden).JSON(dto.StandardResponse{
				Code:    code,
				Message: code.ToMessage("Invalid Auth Key"),
			})
		}

		return c.Next()
	}
}
