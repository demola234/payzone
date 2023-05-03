package token

type Payload struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	IssuedAt int64  `json:"issued_at"`
	ExpiresAt  int64  `json:"expires"`
}