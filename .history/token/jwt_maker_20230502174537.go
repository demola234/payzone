package token


type JWTMaker struct {
	secretKey string
}

func NewJWT