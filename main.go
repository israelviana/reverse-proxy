package main

import (
	"github.com/gofiber/fiber/v2"
	"reverse-proxy/internal/server"
)

func main() {
	app := fiber.New()
	server.InitService(app)
}
