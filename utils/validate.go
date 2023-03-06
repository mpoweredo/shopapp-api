package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ErrorResponse struct {
	Field   string
	Message string
}

var validate = validator.New()

func ValidateStruct[T any](payload T) []*ErrorResponse {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.Field = err.Field()
			element.Message = err.Translate(trans)

			errors = append(errors, &element)
		}
	}
	return errors
}
