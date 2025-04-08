package pvztest

import (
	"context"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/repository"
)

const connString = "postgres://postgres:postgres@localhost:5432/pvz"

func NewRepository(t *testing.T) repository.Repository {
	ctx := context.Background()
	repo, err := repository.New(ctx, connString)
	if err != nil {
		t.Fatal(err)
	}
	return repo
}
