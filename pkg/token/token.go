package token

import (
	"encoding/base64"
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
	str := base64.StdEncoding.EncodeToString(bytes)
	return str, err
}

func (t *jsonParser) Unpack(token string) (Payload, error) {
	bytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return Payload{}, err
	}

	payload := Payload{}
	err = json.Unmarshal(bytes, &payload)

	return payload, err
}
