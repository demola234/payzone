package token

type Payload struct {
	Username string `json:"username"`
	IssuedAt int64  `json:"issued_at"`
	ExpiresAt  int64  `json:"expires"`
}