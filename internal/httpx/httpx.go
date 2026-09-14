package httpx

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func Decode[T any](r *http.Request) (*T, error) {
	var v T
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func writeJson(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, msg string, code string) {
	writeJson(w, status, ErrorResponse{Error: msg, Code: code})
}

func WriteJson(w http.ResponseWriter, status int, payload any) {
	writeJson(w, status, payload)
}
