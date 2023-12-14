package authorization

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.Claims `json:"c,omitempty"`
	AccountID  string `json:"sub"`
	Expiry     int64  `json:"exp"`
	Type       string `json:"typ"`
}

func newClaims(userID string, _type string, expiry int64) *Claims {
	return &Claims{
		AccountID: userID,
		Type:      _type,
		Expiry:    expiry,
	}
}
