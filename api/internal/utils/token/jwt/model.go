package jwt

import (
	"api/internal/utils/token/model"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	model.TokenInfo
}
