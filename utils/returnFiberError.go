package utils

import "github.com/gofiber/fiber/v2"

func ReturnFiberError(c *fiber.Ctx, errorMessage string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": errorMessage,
	})
}
