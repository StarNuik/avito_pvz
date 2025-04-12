package handler

import "github.com/gin-gonic/gin"

func (h *handler) GetPing(ctx *gin.Context) {
	ctx.Status(200)
}
