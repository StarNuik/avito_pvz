package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
)

func (u *usecase) CreateReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error) {
	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockNoWrites)
	if err != nil {
		return entity.Reception{}, nil
	}
	defer tx.Rollback(ctx)

	lastReception, err := u.repo.GetLastReception(ctx, pvzId)
	if err != nil && !errors.Is(err, entity.ErrNotFound) {
		return entity.Reception{}, err
	}
	if err == nil && lastReception.Status == entity.StatusInProgress {
		return entity.Reception{}, entity.ErrAlreadyExists
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
	err = u.repo.CreateReception(ctx, reception)
	if err != nil {
		return entity.Reception{}, err
	}

	return reception, tx.Commit(ctx)
}
