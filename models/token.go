package models

import "github.com/dgrijalva/jwt-go"

type MClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
