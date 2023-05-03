package token

import (
	"errors"
	"time"
	"vendor/golang.org/x/crypto/chacha20poly1305"

	"github.com/o1egl/paseto"
)

// PasetoMaker is a struct that implements Maker interface
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var (
	ErrInvalidKeySize = errors.New("invalid key size")
)

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, ErrInvalidKeySize
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(user string, duration time.Duration) (string, error){
	
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error){

}
