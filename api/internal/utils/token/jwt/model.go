package jwt

import "github.com/golang-jwt/jwt/v5"

type tokenClaims struct {
	jwt.RegisteredClaims
	IP string `json:"ip"`
}
