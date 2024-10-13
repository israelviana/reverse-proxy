package ports

import "github.com/gofiber/fiber/v2"

type IMiddlewareService interface {
	BlockIp(c *fiber.Ctx) error
	RewriteURIMiddleware(c *fiber.Ctx) error
	BlockRequestMiddleware(c *fiber.Ctx) error
}
