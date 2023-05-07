package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type PayloadLogin struct {
	// ID        uuid.UUID `json:"id"`
	EmailName string `json:"emailName"`
	Password  string `json:"password"`
	// IssuedAt  time.Time `json:"issued_at"`
	// ExpiredAt time.Time `json:"expired_at"`
}

type PayloadRegister struct {
	EmailName string `json:"emailName"`
	Password  string `json:"password"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(emailName string) (*PayloadLogin, error) {
	_, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &PayloadLogin{
		// ID:        tokenID,
		EmailName: emailName,
		// IssuedAt:  time.Now(),
		// ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *PayloadLogin) Valid() error {
	if time.Now().After(time.Now().Add(15)) {
		return ErrExpiredToken
	}
	return nil
}
