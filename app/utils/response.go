package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"error":   err.Error(),
	})
}
