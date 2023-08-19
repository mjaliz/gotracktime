package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInavalidToken = errors.New("token is invalid")
	ErrExpiredToken  = errors.New("token is expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID     uuid.UUID `json:"id"`
	UserID uint      `json:"user_id"`
	Email  string    `json:"username"`
	jwt.RegisteredClaims
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:    tokenID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(duration)),
		},
	}
	return payload, nil
}
