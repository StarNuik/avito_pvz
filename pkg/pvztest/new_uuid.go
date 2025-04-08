package pvztest

import (
	"testing"

	"github.com/google/uuid"
)

func NewUuid(t *testing.T) uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	return uuid
}
