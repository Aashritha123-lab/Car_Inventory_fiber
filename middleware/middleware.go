package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SecurityHeader(c *fiber.Ctx) error {

	c.Response().Header.Set("Content-Security-Policy", "default-src 'self'")

	c.Response().Header.Set("X-Frame-Options", "DENY")
	c.Response().Header.Set("X-XSS-Protection", "1; mode=block")
	return c.Next()
}
