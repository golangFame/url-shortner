package auth

import (
	"github.com/goferHiro/url-shortner/internal/genesis"
	"github.com/hiroBzinga/bun"
)
import "github.com/golang-jwt/jwt/v4"

type service struct {
	*genesis.Service
	db *bun.DB
}

type jwtCustomClaims struct {
	Audience string `json:"aud,omitempty"`
	Issuer   string `json:"iss,omitempty"`
	//NotBefore int64  `json:"nbf,omitempty"`
	//Subject   string `json:"sub,omitempty"`

	IssuedAt  int64  `json:"orig_iat,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	ID        string `json:"id,omitempty"`

	*jwt.StandardClaims
}

type MapClaims jwt.MapClaims
