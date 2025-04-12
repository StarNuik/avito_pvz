package token

import (
	"encoding/json"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

type Payload struct {
	UserRole entity.UserRole
}

// TODO doc
type Parser interface {
	Pack(payload Payload) (string, error)
	Unpack(token string) (Payload, error)
}

var _ Parser = (*jsonParser)(nil)

type jsonParser struct{}

func NewParser() Parser {
	return &jsonParser{}
}

// TODO use signed jwt
func (t *jsonParser) Pack(payload Payload) (string, error) {
	bytes, err := json.Marshal(payload)
	return string(bytes), err
}

func (t *jsonParser) Unpack(token string) (Payload, error) {
	payload := Payload{}
	err := json.Unmarshal([]byte(token), &payload)
	return payload, err
}
