package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"shop-app-API/database"
	"shop-app-API/models"
	"strings"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "You are unauthorized to this action"})
	}

	secret := os.Getenv("SECRET")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalidate token"})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid token claim"})
	}

	var user models.User

	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	fmt.Println(user.Id)

	if user.Id != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "The user belonging to this token does not exist"})
	}

	c.Locals("user", user)

	return c.Next()
}
