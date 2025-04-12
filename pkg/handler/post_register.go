package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
)

func (h *handler) PostRegister(ctx *gin.Context) {
	req := dto.PostRegisterJSONBody{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	userRole, err := req.Role.ToEntity()
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	user, err := h.usecase.Register(ctx, string(req.Email), req.Password, userRole)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(201, dto.FromUser(user))
}
