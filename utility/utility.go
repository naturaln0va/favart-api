package utility

import (
	"encoding/json"
	"net/http"
)

// PlainTextMessage defines a text only JSON object
type PlainTextMessage struct {
	Message string `json:"message"`
}

// ErrorMessage defines an error message JSON object
type ErrorMessage struct {
	Error string `json:"error"`
}

// FileInfoMessage defines an object that holds file information
type FileInfoMessage struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// Respond returns JSON and a HTTP status code
func Respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
