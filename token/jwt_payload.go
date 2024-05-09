package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	*Payload
}

// NewJWTPayload creates a new JWT payload
func NewJWTPayload(username string, duration time.Duration) (*JWTPayload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return nil, err
	}
	return &JWTPayload{Payload: payload}, nil
}

///
/// Implement for jwt.Claims interface
///

func (payload *JWTPayload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload *JWTPayload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (payload *JWTPayload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload *JWTPayload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (payload *JWTPayload) GetIssuer() (string, error) {
	return "", nil
}

func (payload *JWTPayload) GetSubject() (string, error) {
	return "", nil
}

// ensure JWTPayload implement jwt.Claims interface
var _ jwt.Claims = (*JWTPayload)(nil)
