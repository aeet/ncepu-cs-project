package oauth2

import (
	"github.com/devcui/ncepu-cs-project/ncu-auth/config"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/golang-jwt/jwt"
)

func NewJWTAccessGenerate() *generates.JWTAccessGenerate {
	option := config.Value.Jwt
	return generates.NewJWTAccessGenerate("", []byte(option.SignedKey), jwt.SigningMethodHS512)
}
