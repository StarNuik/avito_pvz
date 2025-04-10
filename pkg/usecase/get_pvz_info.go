package usecase

import (
	"context"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) GetPvzInfo(ctx context.Context, token token.Payload, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error) {
	if token.UserRole != entity.RoleEmployee && token.UserRole != entity.RoleModerator {
		return entity.PvzInfo{}, entity.ErrUnauthorized
	}

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
