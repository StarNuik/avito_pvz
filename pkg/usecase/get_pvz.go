package usecase

import (
	"context"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) GetPvz(ctx context.Context, token token.Payload, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error) {
	panic("")
}
