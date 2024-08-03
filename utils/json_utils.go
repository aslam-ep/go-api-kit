package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ReadFromRequest reads the JSON request body and dectode it into the provided interface
func ReadFromRequest(w http.ResponseWriter, r *http.Request, requestBody any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return http.ErrNotSupported
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(requestBody); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			http.Error(w, "Request body contains badly formed JSON", http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			http.Error(w, "Request body contains an invalid value for a specifci field", http.StatusBadRequest)
		case errors.Is(err, io.EOF):
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)
		default:
			http.Error(w, "Request body contains invalid JSON", http.StatusBadRequest)
		}
		return err
	}

	return nil
}

// WriteToResponse writes a JSON response with the provided status code and data
func WriteToResponse(w http.ResponseWriter, status int, responseBody any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
