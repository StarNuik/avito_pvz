package usecase

import (
	"context"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (u *usecase) GetPvzInfo(ctx context.Context, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error) {
	if page == nil {
		defaultPage := 1
		page = &defaultPage
	}
	if limit == nil {
		defaultLimit := 10
		limit = &defaultLimit
	}
	offset := (*page - 1) * *limit

	info, err := u.repo.GetPvzInfo(ctx, startDate, endDate, *limit, offset)
	if err != nil {
		return entity.PvzInfo{}, err
	}

	return info, nil
}
