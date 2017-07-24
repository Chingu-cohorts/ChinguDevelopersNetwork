package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// CustomMessage is used to send formatted JSON messages
type CustomMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// JSONMessage simplifies the task of returning JSON data
func JSONMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	var msg = CustomMessage{
		Message: message,
		Code:    code,
	}

	respBody, err := json.MarshalIndent(msg, "", " ")
	if err != nil {
		log.Fatalf("Error formatting JSON message: %v", err)
	}

	w.Write(respBody)
}

// JSONResponse simplifies returning JSON formatted structures
func JSONResponse(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Vary", "Origin; Accept-Encoding")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
