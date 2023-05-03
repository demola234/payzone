package token

import "github.com/o1egl/paseto"

// PasetoMaker is a struct that implements Maker interface
type PasetoMaker struct {
	paseto 	 *paseto.V2
	symmer
}
