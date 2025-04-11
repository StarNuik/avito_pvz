package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/gen"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/repository"
)

type Handler interface {
	Ping(ctx *gin.Context)
}

var _ Handler = (*handler)(nil)

type handler struct {
	repo   repository.Repository
	hasher password.Hasher
	gen    gen.Gen
}

func (h *handler) Ping(ctx *gin.Context) {
	ctx.Status(200)
}

func New(repo repository.Repository, hasher password.Hasher, gen gen.Gen) Handler {
	return &handler{
		repo:   repo,
		hasher: hasher,
		gen:    gen,
	}
}
