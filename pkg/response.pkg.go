package pkg

import "github.com/gofiber/fiber/v2"

func ErrRosponse(c *fiber.Ctx, statusCode int, message string, err string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "failed",
		"message": message,
		"errors":  err,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
