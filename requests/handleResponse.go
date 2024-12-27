package requests

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Message string   `json:"message"`
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

func HandlerResponse(w http.ResponseWriter, code int, payload interface{}) error {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := encoder.Encode(payload); err != nil {
		return fmt.Errorf("error writing JSON: %w", err)
	}
	return nil
}

func HandlerError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Error: %s", message)
	}
	HandlerResponse(w, code, ErrResponse{Error: message})
}