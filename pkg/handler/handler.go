package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
)

type Handler interface {
	GetPing(*gin.Context)
	PostDummyLogin(*gin.Context)
	PostLogin(*gin.Context)
	PostRegister(*gin.Context)
	PostPvz(*gin.Context)
	GetPvz(*gin.Context)
	PostCloseLastReception(*gin.Context)
	PostDeleteLastProduct(*gin.Context)
	PostReceptions(*gin.Context)
	PostProducts(*gin.Context)
}

var _ Handler = (*handler)(nil)

type handler struct {
	usecase     usecase.Usecase
	tokenParser token.Parser
}

func New(usecase usecase.Usecase, tokenParser token.Parser) Handler {
	return &handler{
		usecase,
		tokenParser,
	}
}

// GetPvz implements Handler.
func (h *handler) GetPvz(*gin.Context) {
	panic("unimplemented")
}

// PostCloseLastReception implements Handler.
func (h *handler) PostCloseLastReception(*gin.Context) {
	panic("unimplemented")
}

// PostDeleteLastProduct implements Handler.
func (h *handler) PostDeleteLastProduct(*gin.Context) {
	panic("unimplemented")
}

// PostLogin implements Handler.
func (h *handler) PostLogin(*gin.Context) {
	panic("unimplemented")
}

// PostProducts implements Handler.
func (h *handler) PostProducts(*gin.Context) {
	panic("unimplemented")
}

// PostPvz implements Handler.
func (h *handler) PostPvz(*gin.Context) {
	panic("unimplemented")
}

// PostReceptions implements Handler.
func (h *handler) PostReceptions(*gin.Context) {
	panic("unimplemented")
}
