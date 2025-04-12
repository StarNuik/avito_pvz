package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
)

type Handler interface {
	GetPing(*gin.Context)
	PostDummyLogin(*gin.Context)
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
