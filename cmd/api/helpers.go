package main

import (
	"encoding/json"
	"net/http"
)

// Define an clientResponse type.
type clientResponse map[string]interface{}

// writeJSON is a helper that writes the response to the client in JSON format.
func (app *application) writeJSON(w http.ResponseWriter, status int, data clientResponse, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
