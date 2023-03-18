package router

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/handlers/auth"
	"shop-app-API/handlers/user"
	"shop-app-API/middleware"
)

func SetupRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/signin", auth.Signin)
	authGroup.Post("/signup", auth.Signup)

	userGroup := app.Group("/profile")
	userGroup.Post("/details", middleware.DeserializeUser, user.UpdateProfileDetails)
}
