package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type AccessToken struct {
	jwt.RegisteredClaims
}
