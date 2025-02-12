package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Errors string `json:"errors"`
}

func JSONError(w http.ResponseWriter, err any, code int) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
