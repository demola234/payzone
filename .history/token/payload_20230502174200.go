package token

import (
	"github.com/google/uuid"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiresAt int64     `json:"expires"`
}
