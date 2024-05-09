package token

import "time"

type PasetoPayload struct {
	*Payload
}

// NewPasetoPayload creates a new Paseto payload
func NewPasetoPayload(username string, duration time.Duration) (*PasetoPayload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return nil, err
	}
	return &PasetoPayload{Payload: payload}, nil
}
