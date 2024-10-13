package services

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"reverse-proxy/internal/core/ports"
)

var _ ports.IReverseProxyService = (*ReverseProxyService)(nil)

type ReverseProxyService struct{}

func NewReverseProxyService() *ReverseProxyService {
	return &ReverseProxyService{}
}

func (r ReverseProxyService) RedirectRequest(c *fiber.Ctx) error {
	backendURL := "http://backend-server.com" + c.OriginalURL()

	reqBody := c.Body()

	req, err := http.NewRequest(c.Method(), backendURL, bytes.NewReader(reqBody))
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	c.Request().Header.VisitAll(func(key, value []byte) {
		req.Header.Add(string(key), string(value))
	})

	req.Header.Add("X-Forwarded-For", c.IP())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(resp.StatusCode).Send(body)
}
