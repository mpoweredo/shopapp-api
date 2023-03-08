package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"shop-app-API/models"
	"strconv"
	"time"
)

func CreateJWT(user models.User) (string, int64, error) {
	jwtMaxAge, err := strconv.Atoi(os.Getenv("JWT_MAX_AGE"))

	if err != nil {
		return "", 0, errors.New("couldn't parse jwtMaxAge")
	}

	exp := time.Now().Add(time.Minute * time.Duration(jwtMaxAge)).Unix()
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.Id
	claims["exp"] = exp

	secret := os.Getenv("SECRET")

	t, err := tokenByte.SignedString([]byte(secret))

	if err != nil {
		return "", 0, errors.New("generating JWT failed")
	}

	return t, exp, nil
}
