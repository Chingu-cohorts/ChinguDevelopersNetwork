package models

// ResponseToken represents a signed JWT token
type ResponseToken struct {
	Token string `json:"token"`
}
