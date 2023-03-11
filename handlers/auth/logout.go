package auth

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/resources"
	"time"
)

func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"snackbar": resources.SnackbarResponse{Message: "Successfully signed out!", Type: resources.SUCCESS},
	})
}
