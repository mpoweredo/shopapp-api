package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"strings"
	"unicode"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateStruct[T any](payload T) fiber.Map {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	var errors []*ErrorResponse
	err := validate.Struct(payload)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			runes := []rune(err.Field())

			runes[0] = unicode.ToLower(runes[0])
			element.Field = string(runes)
			element.Message = convertCamelCase(err.Translate(trans))

			if err.Field() == "PasswordConfirm" || err.Field() == "Password" {
				if err.Translate(trans) == "Password must be equal to PasswordConfirm" || err.Translate(trans) == "PasswordConfirm must be equal to Password" {
					element.Message = "Passwords do not match"
				}
			}

			errors = append(errors, &element)
		}
	}
	if len(errors) > 0 {
		return fiber.Map{"errorFields": errors}
	}
	return nil
}

func convertCamelCase(input string) string {
	var output strings.Builder
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 && unicode.IsLower(rune(input[i-1])) {
				output.WriteString(" ")
			}
			output.WriteRune(unicode.ToLower(r))
		} else {
			output.WriteRune(r)
		}
	}

	s := output.String()

	return strings.ToUpper(s[0:1]) + s[1:]
}
