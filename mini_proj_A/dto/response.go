package dto

import (
	"encoding/json"
	"net/http"
)

// APIResponse is the standardized JSON envelope returned to clients.
type APIResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    any               `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// SendJSON helper sets Content-Type header and encodes response.
func SendJSON(w http.ResponseWriter, statusCode int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

// SendSuccess sends a 2xx standardized response with data.
func SendSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	SendJSON(w, statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendError sends an error response with message and optional field-level error details.
func SendError(w http.ResponseWriter, statusCode int, message string, fieldErrors map[string]string) {
	SendJSON(w, statusCode, APIResponse{
		Success: false,
		Message: message,
		Errors:  fieldErrors,
	})
}
