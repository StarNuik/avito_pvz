package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
)

func (u *usecase) CloseLastReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error) {
	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockNoWrites)
	if err != nil {
		return entity.Reception{}, err
	}
	defer tx.Rollback(ctx)

	reception, err := u.repo.GetLastReception(ctx, pvzId)
	if err != nil {
		return entity.Reception{}, err
	}

	if reception.Status != entity.StatusInProgress {
		return entity.Reception{}, entity.ErrReceptionClosed
	}

	reception, err = u.repo.UpdateReceptionStatus(ctx, reception.Id, entity.StatusClosed)
	if err != nil {
		return entity.Reception{}, err
	}

	return reception, tx.Commit(ctx)
}
