package token

import "time"

// Maker is an interface that creates and verifies tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(user string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
