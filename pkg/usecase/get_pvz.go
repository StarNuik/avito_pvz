package usecase

import (
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
	"time"
)

func (u *usecase) GetPvz(token token.Payload, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error) {
	panic("")
}
