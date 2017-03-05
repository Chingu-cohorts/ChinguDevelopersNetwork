package models

// CustomMessage helps returning valid JSON data
type CustomMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
