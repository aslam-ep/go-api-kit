package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ReadFromRequest reads the JSON request body and dectode it into the provided interface
func ReadFromRequest(r *http.Request, requestBody any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("content-type header is not application/json")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(requestBody); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return errors.New("request body contains badly formed JSON")
		case errors.As(err, &unmarshalTypeError):
			return errors.New("request body contains an invalid value for a specifci field")
		case errors.Is(err, io.EOF):
			return errors.New("request body must not be empty")
		default:
			return errors.New("request body contains invalid JSON")
		}
	}

	return nil
}

// WriterErrorResponse writes a Error JSOM response with the provided status code and error message
func WriterErrorResponse(w http.ResponseWriter, status int, message string) {
	res := &MessageRes{
		Success: false,
		Message: message,
	}

	WriteResponse(w, status, res)
}

// WriteToResponse writes a JSON response with the provided status code and data
func WriteResponse(w http.ResponseWriter, status int, responseBody any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
