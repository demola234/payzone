package token

import "time"

type Maker interface {
	CreateToken(user string, duration time.Duration) (String)
}