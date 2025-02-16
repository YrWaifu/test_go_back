package errors

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Errors string `json:"errors"`
}

func JSONError(w http.ResponseWriter, err any, code int) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	encodeErr := json.NewEncoder(w).Encode(err)
	if encodeErr != nil {
		slog.Info("encode")
	}
}
