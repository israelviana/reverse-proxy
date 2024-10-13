package server

import (
	"github.com/gofiber/fiber/v2"
	"reverse-proxy/internal/core/services"
	"reverse-proxy/internal/repositories"
)

func InitService(app *fiber.App) {
	postgres := repositories.NewPostgreSQL()
	err := postgres.StartConnection()
	if err != nil {
		panic(err)
	}

	middlewareService := services.NewMiddlewareService(postgres)

	app.Use(middlewareService.BlockIp)
	app.Use(middlewareService.RewriteURIMiddleware)
	app.Use(middlewareService.BlockRequestMiddleware)

	reverseProxyService := services.NewReverseProxyService()

	app.All("/*", reverseProxyService.RedirectRequest)

	app.Listen("0.0.0.0:3000")
}
