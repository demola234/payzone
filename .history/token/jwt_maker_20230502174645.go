package token

const minSecretKey = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len (secretKey) < minSecretKey {
		return nil, ErrInvalidSecretKey
	}
}
