package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"shop-app-API/database"
	"shop-app-API/models"
	"shop-app-API/utils"
	"strconv"
	"strings"
	"time"
)

type SignUpInput struct {
	Username        string `json:"username" validate:"required,min=5,max=24"`
	Email           string `json:"email" validate:"required,email,min=6,max=48"`
	Password        string `json:"password" validate:"required,min=5,max=24"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=5,max=24"`
}

func Signup(c *fiber.Ctx) error {

	b := SignUpInput{}

	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(b)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if b.Password != b.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	fmt.Println(string(hashedPassword))

	user := models.User{
		Username: b.Username,
		Email:    b.Email,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": "Something went wrong"})
	}

	token, exp, err := createJWT(user)

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
		"message": "Added",
		"token":   token,
		"expires": exp,
	})
}

func createJWT(user models.User) (string, int64, error) {
	jwtMaxAge, err := strconv.Atoi(os.Getenv("JWT_MAX_AGE"))

	if err != nil {
		return "", 0, errors.New("couldn't parse jwtMaxAge")
	}

	exp := time.Now().Add(time.Minute * time.Duration(jwtMaxAge)).Unix()
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = exp

	secret := os.Getenv("SECRET")

	t, err := tokenByte.SignedString([]byte(secret))

	if err != nil {
		return "", 0, errors.New("generating JWT failed")
	}

	return t, exp, nil
}
