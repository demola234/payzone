package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Duration     `json:"issued_at"`
	ExpiresAt time.Duration     `json:"expires"`
}
