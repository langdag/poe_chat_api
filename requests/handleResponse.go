package requests

import (
	"encoding/json"
	"log"
	"net/http"
)

type successResponse struct {
	Message string   `json:"message"`
	Data interface{} `json:"data"`
}

type errResponse struct {
	Error string `json:"error"`
}

func HandlerResponse(w http.ResponseWriter, code int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		log.Printf("Response data: %s", data)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HandlerError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("Error: %s", message)
	}
	HandlerResponse(w, code, errResponse{Error: message})
}