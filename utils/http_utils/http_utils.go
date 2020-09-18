package http_utils

import (
	"encoding/json"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors.RestError) {
	RespondJson(w, err.Status(), err)
}
