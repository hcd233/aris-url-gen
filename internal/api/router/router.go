package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/Aris-url-gen/internal/api/handler"
	"github.com/hcd233/Aris-url-gen/internal/api/service"
)

func RegisterRouter(app *fiber.App) {
	healthCheckHandler := handler.NewHealthCheckHandler()
	shortenUrlHandler := handler.NewShortURLHandler(handler.WithShortURLService(service.NewShortenUrlService()))

	healthRouter := app.Group("/health")
	healthRouter.Get("", healthCheckHandler.HealthCheck)

	v1Router := app.Group("/v1")
	shortURLRouter := v1Router.Group("/shortURL")

	shortURLRouter.Post("", shortenUrlHandler.GenerateShortURL)
	shortURLRouter.Get("/:shortURL", shortenUrlHandler.GetOriginalURL)
}
