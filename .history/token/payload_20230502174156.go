package token

import "github.com/google/uuid"

type Payload struct {
	ID uuid. `json:"id"`
	Username string `json:"username"`
	IssuedAt int64  `json:"issued_at"`
	ExpiresAt  int64  `json:"expires"`
}