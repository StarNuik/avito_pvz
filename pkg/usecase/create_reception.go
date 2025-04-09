package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreateReception(ctx context.Context, token token.Payload, pvzId uuid.UUID) (entity.Reception, error) {
	if token.UserRole != entity.RoleEmployee {
		return entity.Reception{}, entity.ErrUnauthorized
	}

	id, err := u.gen.Uuid()
	if err != nil {
		return entity.Reception{}, err
	}

	now := u.gen.Now()

	err = u.repo.LockPvz(ctx, pvzId, repository.LockAllowWrites)
	if err != nil {
		return entity.Reception{}, nil
	}

	reception := entity.Reception{
		Id:       id,
		PvzId:    pvzId,
		DateTime: now,
		Status:   entity.StatusInProgress,
	}
	reception, err = u.repo.CreateReception(ctx, reception)

	return reception, err
}
