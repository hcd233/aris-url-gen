package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/handler"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
)

func RegisterRouter(app *fiber.App) {
	healthCheckHandler := handler.NewHealthCheckHandler()
	shortURLHandler := handler.NewShortURLHandler(handler.WithShortURLService(service.NewShortURLService()))

	healthRouter := app.Group("/health")
	healthRouter.Get("", healthCheckHandler.HealthCheck)

	v1Router := app.Group("/v1")
	shortURLRouter := v1Router.Group("/shortURL")

	shortURLRouter.Post("", shortURLHandler.GenerateShortURL)
	shortURLRouter.Get("/:shortURL", shortURLHandler.GetOriginalURL)
}
