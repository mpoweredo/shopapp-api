package router

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/handlers/auth"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/signup", handlers.Signup)
	auth.Post("/signin", handlers.Signin)
}
