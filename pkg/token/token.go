package token

import "github.com/starnuik/avito_pvz/pkg/entity"

type Token interface {
	String() string
}

type Payload struct {
	UserRole entity.UserRole
}

func Pack(p Payload) Token {
	panic("")
}

func Unpack(token string) Payload {
	panic("")
}
