package gen

//go:generate mockgen -destination=../mocks/mock_gen.go -package=mocks github.com/starnuik/avito_pvz/pkg/gen Gen

import (
	"time"

	"github.com/google/uuid"
)

// TODO doc
type Gen interface {
	Now() time.Time
	Uuid() (uuid.UUID, error)
}

var _ Gen = (*impl)(nil)

type impl struct{}

func New() Gen {
	return &impl{}
}

func (*impl) Now() time.Time {
	return time.Now().UTC()
}

func (*impl) Uuid() (uuid.UUID, error) {
	return uuid.NewRandom()
}
