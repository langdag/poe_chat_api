package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(payload); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, payload interface{}) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(payload); err != nil {
		return fmt.Errorf("error writing JSON: %w", err)
	}
	return nil
}
