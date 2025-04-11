package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/gen"
	"github.com/starnuik/avito_pvz/pkg/handler"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/repository"
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

	repo, err := repository.New(context.TODO(), "postgres://postgres:postgres@localhost:5432/pvz")
	if err != nil {
		return nil, err
	}

	handler := handler.New(repo, hasher, gen)

	router := gin.Default()

	router.GET("/ping", handler.Ping)

	return &app{
		Engine: router,
	}, nil
}

func (a *app) Run() {
	// TODO graceful shutdown
	a.Engine.Run(":8080")
}
