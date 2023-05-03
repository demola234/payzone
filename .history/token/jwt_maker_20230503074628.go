package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKey = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKey)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(user string, duration time.Duration) (string, error) {
	payload, err := NewPayLoad(user, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyfunc := func(token jwt.Token) ({}interface, error) {
	
	}
	jwt.ParseWithClaims(token, &Payload{}, keyfunc)
}
