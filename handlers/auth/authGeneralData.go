package auth

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/utils"
)

func GetAuthGeneralData(c *fiber.Ctx) error {
	user := utils.GetUser(c)

	return c.Status(200).JSON(fiber.Map{
		"data": fiber.Map{
			"id":       user.Id,
			"email":    user.Email,
			"username": user.Username,
		},
	})
}
