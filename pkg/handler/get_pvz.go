package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/dto"
)

func (h *handler) GetPvz(ctx *gin.Context) {

	startDate, err := getTimeQuery(ctx, "startDate")
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	endDate, err := getTimeQuery(ctx, "endDate")
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	var page, limit *int
	if val, err := getIntQuery(ctx, "page"); err == nil {
		page = &val
	}
	if val, err := getIntQuery(ctx, "limit"); err == nil {
		limit = &val
	}

	info, err := h.usecase.GetPvzInfo(ctx, startDate, endDate, page, limit)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.JSON(200, dto.FromPvzInfo(info))
}

func getTimeQuery(ctx *gin.Context, key string) (time.Time, error) {
	str := ctx.Query(key)
	val, err := time.Parse(time.RFC3339, str)
	return val, err
}

func getIntQuery(ctx *gin.Context, key string) (int, error) {
	str := ctx.Query(key)
	val, err := strconv.Atoi(str)
	return val, err
}
