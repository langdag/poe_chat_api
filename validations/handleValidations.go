package validations

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"encoding/json"
	"github.com/langdag/poe_chat_api/requests"
)

// Struct to hold error messages for each field and tag
type FieldValidationMessages struct {
	Required string `json:"required"`
	Format    string `json:"format"`
}

// Struct to hold all field validation messages
type UserValidationMessages struct {
	Username FieldValidationMessages `json:"Username"`
	Email    FieldValidationMessages `json:"Email"`
	Password FieldValidationMessages `json:"Password"`
}

// Example of how to use the struct
var CustomErrorMessages = UserValidationMessages{
	Username: FieldValidationMessages{
		Required: "Username cannot be empty.",
	},
	Email: FieldValidationMessages{
		Required: "Email is required.",
		Format:    "Invalid email format.",
	},
	Password: FieldValidationMessages{
		Required: "Password cannot be empty.",
	},
}

// HandleValidations validates the struct and sends custom error messages
func HandleValidations(w http.ResponseWriter, obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			customErrors := parseValidationErrors(validationErrors)

			// Send JSON response with errors as an object
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errors": customErrors,
			})
			return err
		}

		// Fallback for non-validation errors
		requests.HandlerError(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}

// parseValidationErrors parses the validation errors and retrieves custom messages
func parseValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, fieldErr := range validationErrors {
		fieldName := fieldErr.Field()
		tag := fieldErr.Tag()

		// Map the field name to the correct struct and retrieve the error message
		switch fieldName {
		case "Username":
			if tag == "required" {
				errors[fieldName] = CustomErrorMessages.Username.Required
			}
		case "Email":
			if tag == "required" {
				errors[fieldName] = CustomErrorMessages.Email.Required
			} else if tag == "email" {
				errors[fieldName] = CustomErrorMessages.Email.Format
			}
		case "Password":
			if tag == "required" {
				errors[fieldName] = CustomErrorMessages.Password.Required
			}
		default:
			// Default error message for unrecognized fields
			errors[fieldName] = "Validation failed for " + fieldName
		}
	}
	return errors
}
