package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/hcd233/Aris-url-gen/docs"
	"github.com/hcd233/Aris-url-gen/internal/api/handler"
	auth "github.com/hcd233/Aris-url-gen/internal/api/middleware"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
	"github.com/hcd233/Aris-url-gen/internal/config"
)

// RegisterRouter 注册路由
//
//	param app *fiber.App
//	author centonhuang
//	update 2025-01-25 14:05:06
func RegisterRouter(app *fiber.App) {
	healthCheckHandler := handler.NewHealthCheckHandler()
	shortURLHandler := handler.NewShortURLHandler(handler.WithShortURLService(service.NewShortURLService()))

	if config.APIMode != config.ModeProd {
		// swagger
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	healthRouter := app.Group("/health")
	healthRouter.Get("", healthCheckHandler.HealthCheck)

	v1Router := app.Group("/v1")
	v1Router.Post("/shortURL", auth.New(), shortURLHandler.GenerateShortURL)
	v1Router.Get("/s/:shortURL", shortURLHandler.GetOriginalURL)
}
