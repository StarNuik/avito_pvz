package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreatePvz(ctx context.Context, token token.Payload, city entity.PvzCity, id *uuid.UUID, registrationDate *time.Time) (entity.Pvz, error) {
	panic("")
}
