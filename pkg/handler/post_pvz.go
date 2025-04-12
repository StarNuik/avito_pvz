package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (h *handler) PostPvz(ctx *gin.Context) {
	tokenPayload, ok := GetTokenPayload(ctx)
	if !ok {
		ctx.AbortWithError(403, fmt.Errorf("incorrect token"))
		return
	}

	req := dto.PVZ{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	city, err := req.City.ToEntity()
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	pvz, err := h.usecase.CreatePvz(ctx, tokenPayload, city, req.Id, req.RegistrationDate)
	if errors.Is(err, entity.ErrUnauthorized) {
		ctx.AbortWithError(403, err)
		return
	}

	ctx.JSON(201, dto.FromPvz(pvz))
}
