package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreateReception(ctx context.Context, token token.Payload, pvzId uuid.UUID) (entity.Reception, error) {
	if token.UserRole != entity.RoleEmployee {
		return entity.Reception{}, entity.ErrUnauthorized
	}

	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockNoWrites)
	if err != nil {
		return entity.Reception{}, nil
	}
	defer tx.Rollback()

	_, err = u.repo.GetOpenReception(ctx, pvzId)
	if err == nil {
		return entity.Reception{}, entity.ErrAlreadyExists
	}
	if !errors.Is(err, entity.ErrNotFound) {
		return entity.Reception{}, err
	}

	id, err := u.gen.Uuid()
	if err != nil {
		return entity.Reception{}, err
	}

	now := u.gen.Now()

	reception := entity.Reception{
		Id:       id,
		PvzId:    pvzId,
		DateTime: now,
		Status:   entity.StatusInProgress,
	}

	reception, err = u.repo.CreateReception(ctx, reception)
	if err != nil {
		return entity.Reception{}, err
	}

	return reception, tx.Commit()
}
