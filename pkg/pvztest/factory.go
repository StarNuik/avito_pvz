package pvztest

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/app"
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

func NewUuid(t *testing.T) uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	return uuid
}

func NewApp(t *testing.T) app.App {
	app, err := app.New()
	if err != nil {
		t.Fatal(err)
	}
	return app
}
