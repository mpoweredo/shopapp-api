package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/utils"
)

type SigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signin(c *fiber.Ctx) error {

	b := SigninInput{}

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User

	result := database.DB.First(&user, "email = ?", b.Email)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Credentials are invalid"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(b.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Credentials are invalid"})
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
		"message": "Successfully signed in!",
		"token":   token,
		"expires": exp,
	})

}
