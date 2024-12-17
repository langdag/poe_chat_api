package validations

import (
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/langdag/poe_chat_api/requests"
)

func HandleValidations(w http.ResponseWriter, obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err != nil {
		requests.HandlerError(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return err
	}

	return nil
}

