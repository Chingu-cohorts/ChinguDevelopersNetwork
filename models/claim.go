package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim contains the data of the token
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
