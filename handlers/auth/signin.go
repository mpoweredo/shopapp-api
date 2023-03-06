package handlers

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/models"
)

func Signin(c *fiber.Ctx) error {

	b := models.User{}

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil

}
