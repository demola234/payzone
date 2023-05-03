package token

import "fmt"

const minSecretKey = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len (secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKey)
	}
	return &JWTMaker{secretKey}, nil
}


CreateToken(user string, duration time.Duration) (string, error)

// VerifyToken checks if the token is valid or not
VerifyToken(token string) (*Payload, error)