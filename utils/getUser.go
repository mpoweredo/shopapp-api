package utils

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/models"
)

func GetUser(c *fiber.Ctx) models.User {
	user := c.Locals("profile").(models.User)

	return user
}
