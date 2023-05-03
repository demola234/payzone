package token

type Maker interface {
	CreateToken(user string, duration )
}