package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

// Custom error messages map
var customErrorMessages = map[string]string{
	"required": "The {field} field is required.",
	"email":    "The {field} field must be a valid email address.",
	"min":      "The {field} field must be at least {param} characters long.",
}

// GetValidator returns the validator instance, initializing it if necessary
func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		registerCustomMessages(validate)
	}
	return validate
}

// Register custom error messages
func registerCustomMessages(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field validator.FieldError) string {
		return field.Field()
	})

	for tag, message := range customErrorMessages {
		tag, message := tag, message
		_ = validate.RegisterTranslation(tag, validate.Translations(), func(ut validator.TranslationRegisterFunc) error {
			return ut.Add(tag, message, true)
		}, func(ut validator.TranslationFunc) error {
			return nil
		})
	}
}

// ValidateStruct validates a struct and returns a map of field errors with custom messages
func ValidateStruct(s any) error {
	err := GetValidator().Struct(s)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errorMessages := make(map[string]string)

	for _, e := range validationErrors {
		fieldName := e.Field()
		tag := e.Tag()
		message := customErrorMessages[tag]

		// Replace {field} and {param} placeholders in the message
		message = fmt.Sprintf(message, fieldName, e.Param())

		errorMessages[fieldName] = message
	}

	return formatErrors(errorMessages)
}

// Format the error messages into a single error
func formatErrors(errs map[string]string) error {
	var errorMessage string
	for field, msg := range errs {
		errorMessage += fmt.Sprintf("%s: %s\n", field, msg)
	}
	return errors.New(errorMessage)
}
