package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Exp      int64  `json:"exp"`
}

func (m MyClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(m.Exp, 0)), nil
}

func (m MyClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return nil, errors.New("issued_at not set in claims")
}

func (m MyClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, errors.New("not_before not supported in claims")
}

func (m MyClaims) GetIssuer() (string, error) {
	return "", errors.New("issuer not set in claims")
}

func (m MyClaims) GetSubject() (string, error) {
	return "", errors.New("subject not commonly used for user information")
}

func (m MyClaims) GetAudience() (jwt.ClaimStrings, error) {
	return nil, errors.New("audience not set in claims")
}
