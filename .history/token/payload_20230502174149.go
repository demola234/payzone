package token

type Payload struct {
	ID 
	Username string `json:"username"`
	IssuedAt int64  `json:"issued_at"`
	ExpiresAt  int64  `json:"expires"`
}