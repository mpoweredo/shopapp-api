package router

import (
	"github.com/gofiber/fiber/v2"
	"shop-app-API/handlers/auth"
	"shop-app-API/handlers/profile"
	"shop-app-API/middleware"
)

func SetupRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/signin", auth.Signin)
	authGroup.Post("/signup", auth.Signup)
	authGroup.Post("/logout", auth.Logout)

	userGroup := app.Group("/profile")
	userGroup.Post("/details", middleware.DeserializeUser, profile.UpdateProfileDetails)

	userGroup.Post("/delivery", middleware.DeserializeUser, profile.AddDeliveryAddress)
	userGroup.Get("/delivery", middleware.DeserializeUser, profile.GetDeliveryAddresses)
	userGroup.Delete("/delivery/:id", middleware.DeserializeUser, profile.DeleteDeliveryAddress)
	userGroup.Put("/delivery/:id", middleware.DeserializeUser, profile.UpdateDeliveryAddress)
}
