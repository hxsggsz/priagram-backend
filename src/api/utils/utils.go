package utils

import (
	"encoding/json"
	"net/http"
)

type apiReturn struct {
	Status  int
	Message string
}

func newApiReturn(status int, message string) apiReturn {
	return apiReturn{
		Status: status, Message: message,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, newApiReturn(status, err.Error()))
}
