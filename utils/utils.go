package utils

import (
	"encoding/json"
	"net/http"

	"github.com/Oxyrus/ChinguCentral/models"
)

// ErrorWithJSON simplifies returning a JSON error
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	var msg = models.CustomMessage{
		Code:    code,
		Message: message,
	}

	respBody, err := json.MarshalIndent(msg, "", " ")
	if err != nil {
		panic(err)
	}

	w.Write(respBody)
}

// ResponseWithJSON simplifies returning a successful response
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
