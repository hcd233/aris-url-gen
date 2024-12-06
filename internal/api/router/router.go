package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/handler"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
)

// RegisterRouter 注册路由
//
//	@param app *fiber.App
//	@author centonhuang
//	@update 2024-12-05 16:13:09
func RegisterRouter(app *fiber.App) {
	healthCheckHandler := handler.NewHealthCheckHandler()
	shortURLHandler := handler.NewShortURLHandler(handler.WithShortURLService(service.NewShortURLService()))

	healthRouter := app.Group("/health")
	healthRouter.Get("", healthCheckHandler.HealthCheck)

	v1Router := app.Group("/v1")
	v1Router.Post("/shortURL", shortURLHandler.GenerateShortURL)
	v1Router.Get("/s/:shortURL", shortURLHandler.GetOriginalURL)
}
