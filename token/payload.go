package token

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

///
/// Implement for jwt.Claims interface
///

// GetExpirationTime implements the Claims interface.
func (c *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(c.ExpiredAt), nil
}

// GetNotBefore implements the Claims interface.
func (c *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

// GetIssuedAt implements the Claims interface.
func (c *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(c.IssuedAt), nil
}

// GetAudience implements the Claims interface.
func (c *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

// GetIssuer implements the Claims interface.
func (c *Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetSubject implements the Claims interface.
func (c *Payload) GetSubject() (string, error) {
	return "", nil
}
