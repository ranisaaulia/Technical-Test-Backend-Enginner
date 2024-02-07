package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseJSON(c *fiber.Ctx, p interface{}, status int) error {
	return c.Status(status).JSON(p)
}
