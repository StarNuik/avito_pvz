package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreatePvz(ctx context.Context, token token.Payload, city entity.PvzCity, id *uuid.UUID, registrationDate *time.Time) (entity.Pvz, error) {
	if token.UserRole != entity.RoleModerator {
		return entity.Pvz{}, entity.ErrUnauthorized
	}

	if id == nil {
		newId, err := u.gen.Uuid()
		if err != nil {
			return entity.Pvz{}, err
		}
		id = &newId
	}

	if registrationDate == nil {
		now := u.gen.Now()
		registrationDate = &now
	}

	pvz := entity.Pvz{
		Id:               *id,
		RegistrationDate: *registrationDate,
		City:             city,
	}
	err := u.repo.CreatePvz(ctx, pvz)

	return pvz, err
}
