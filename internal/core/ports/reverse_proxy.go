package ports

import "github.com/gofiber/fiber/v2"

type IReverseProxyService interface {
	RedirectRequest(c *fiber.Ctx) error
}
