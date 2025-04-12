package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (h *handler) PostPvz(ctx *gin.Context) {
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

	pvz, err := h.usecase.CreatePvz(ctx, city, req.Id, req.RegistrationDate)
	if errors.Is(err, entity.ErrUnauthorized) {
		ctx.AbortWithError(403, err)
		return
	}

	ctx.JSON(201, dto.FromPvz(pvz))
}
