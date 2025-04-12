package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

// PostLogin implements Handler.
func (h *handler) PostLogin(ctx *gin.Context) {
	req := dto.PostLoginJSONBody{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	payload, err := h.usecase.Login(ctx, string(req.Email), req.Password)
	if errors.Is(err, entity.ErrIncorrectLogin) {
		ctx.AbortWithError(401, err)
		return
	}
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	token, err := h.tokenParser.Pack(payload)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(200, token)
}
