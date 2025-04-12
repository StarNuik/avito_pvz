package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
)

func (h *handler) PostDummyLogin(ctx *gin.Context) {
	dto := dto.PostDummyLoginJSONBody{}
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	userRole, err := dto.Role.ToEntity()
	if err != nil {
		ctx.AbortWithError(400, err)
	}

	payload := h.usecase.DummyLogin(userRole)

	token, err := h.tokenParser.Pack(payload)
	if err != nil {
		ctx.AbortWithError(500, err)
	}

	ctx.JSON(200, token)
}
