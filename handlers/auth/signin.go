package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/resources"
	"shop-app-API/utils"
)

type SigninRequest struct {
	Email    string `json:"email" validate:"required,email,min=6,max=48"`
	Password string `json:"password"`
}

func Signin(c *fiber.Ctx) error {

	b := SigninRequest{}

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(b)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user models.User

	result := database.DB.First(&user, "email = ?", b.Email)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"snackbar": resources.SnackbarResponse{Message: "Credentials are invalid!", Type: resources.ERROR},
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(b.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"snackbar": resources.SnackbarResponse{Message: "Credentials are invalid!", Type: resources.ERROR},
		})
	}

	fmt.Println(user)

	token, exp, err := utils.CreateJWT(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing JWT",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   int(exp),
	})

	return c.Status(200).JSON(fiber.Map{
		"snackbar": resources.SnackbarResponse{Message: "Successfully signed in!", Type: resources.SUCCESS},
		"token":    token,
		"expires":  exp,
	})

}
