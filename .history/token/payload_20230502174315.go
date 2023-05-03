package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires"`
}


func NewPayLoad(username string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Username: username,
		IssuedAt: time.Now(),
	}

	if duration > 0 {
		payload.ExpiresAt = payload.IssuedAt.Add(duration)
	}

	payload.ID = uuid.New()

	return payload, nil
}