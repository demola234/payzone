package token

import (
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.     `json:"issued_at"`
	ExpiresAt int64     `json:"expires"`
}
