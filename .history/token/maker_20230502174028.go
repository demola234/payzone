package token

import "time"

// Maker is an interface that creates and verifies tokens
type Maker interface {
	// 
	CreateToken(user string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
