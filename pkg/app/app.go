package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/api"
	"github.com/starnuik/avito_pvz/pkg/gen"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/usecase"
)

type App interface {
	Run()
}

type app struct {
	*gin.Engine
}

func New() (App, error) {
	hasher := password.NewHasher()
	gen := gen.New()

	// TODO context, conn string
	repo, err := repository.New(context.TODO(), "postgres://postgres:postgres@localhost:5432/pvz")
	if err != nil {
		return nil, err
	}

	usecase := usecase.New(repo, hasher, gen)

	handler := api.New(usecase)

	router := gin.Default()

	api.RegisterHandlers(router, handler)

	return &app{
		Engine: router,
	}, nil
}

func (a *app) Run() {
	// TODO graceful shutdown
	a.Engine.Run(":8080")
}
