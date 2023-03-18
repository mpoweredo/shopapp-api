package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/resources"
	"shop-app-API/utils"
	"strings"
)

type SignupRequest struct {
	Username        string `json:"username" validate:"required,min=5,max=24"`
	Email           string `json:"email" validate:"required,email,min=6,max=48"`
	Password        string `json:"password" validate:"required,min=5,max=24,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"passwordConfirm" validate:"eqfield=Password"`
}

func Signup(c *fiber.Ctx) error {

	b := SignupRequest{}

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(b)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Something went wrong while encrypting password"})
	}

	user := models.User{
		Username: b.Username,
		Email:    b.Email,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"snackbar": resources.SnackbarResponse{Message: "Account with this email already exists!", Type: resources.ERROR},
		})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Something went wrong"})
	}

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
		"snackbar": resources.SnackbarResponse{Message: "Successfully signed up!", Type: resources.SUCCESS},
		"token":    token,
		"expires":  exp,
	})
}
