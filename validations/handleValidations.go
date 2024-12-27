package validations

import (
    "github.com/go-playground/validator/v10"
    "net/http"
    "github.com/langdag/poe_chat_api/requests"
)

// Custom error messages for each field and tag
type FieldValidationMessages struct {
    Required string `json:"required,omitempty"`
    Format   string `json:"format,omitempty"`
}

// Struct to hold all field validation messages
type UserValidationMessages struct {
    Username FieldValidationMessages `json:"Username"`
    Email    FieldValidationMessages `json:"Email"`
    Password FieldValidationMessages `json:"Password"`
}

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
        if validationErrors, ok := err.(validator.ValidationErrors); ok {
            customErrors := parseValidationErrors(validationErrors)
            requests.HandlerResponse(w, http.StatusBadRequest, customErrors)
            return err
        }
        requests.HandlerError(w, http.StatusBadRequest, err.Error()) // Assuming this function exists
        return err
    }
    return nil
}

// parseValidationErrors parses the validation errors and retrieves custom messages
func parseValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
    errors := make(map[string]string)
    for _, err := range validationErrors {
        fieldName := err.Field()
        tag := err.Tag()

        if msg, ok := getCustomMessage(fieldName, tag); ok {
            errors[fieldName] = msg
        } else {
            errors[fieldName] = "Validation failed for " + fieldName
        }
    }
    return errors
}

// getCustomMessage retrieves the custom error message for a given field and tag
func getCustomMessage(fieldName, tag string) (string, bool) {
    switch fieldName {
    case "Username":
        if tag == "required" {
            return CustomErrorMessages.Username.Required, true
        }
    case "Email":
        if tag == "required" {
            return CustomErrorMessages.Email.Required, true
        } else if tag == "email" {
            return CustomErrorMessages.Email.Format, true
        }
    case "Password":
        if tag == "required" {
            return CustomErrorMessages.Password.Required, true
        }
    }
    return "", false
}
