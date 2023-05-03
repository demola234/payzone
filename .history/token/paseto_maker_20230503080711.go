package token

import (
	"vendor/golang.org/x/crypto/chacha20poly1305"

	"github.com/o1egl/paseto"
)

// PasetoMaker is a struct that implements Maker interface
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, ErrInvalidKeySize
	}
}
