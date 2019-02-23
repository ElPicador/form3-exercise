package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON marshal & sends `i` as JSON
func WriteJSON(w http.ResponseWriter, statusCode int, i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		Write500(w)
		return
	}
	WriteContentTypeJSON(w, statusCode, bytes)
}

// WriteContentTypeJSON sends a http response with content type json
func WriteContentTypeJSON(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		_, _ = w.Write(body)
	}
}

type statusMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// WriteJSONMessage sends a response formatted in json with statusCode and the message
func WriteJSONMessage(w http.ResponseWriter, statusCode int, message string) {
	body, err := json.Marshal(statusMessage{
		Code:    statusCode,
		Message: message,
	})

	if err != nil {
		log.Printf("[ERROR] cannot marshal json response: %s\n", err)
		Write500(w)
		return
	}
	WriteContentTypeJSON(w, statusCode, body)
}

var jsonBody500 = []byte(`{"status":500, "message": "Internal Server Error"}`)

// Write500 sends a generic 500 error
func Write500(w http.ResponseWriter) {
	WriteContentTypeJSON(w, http.StatusInternalServerError, jsonBody500)
}
