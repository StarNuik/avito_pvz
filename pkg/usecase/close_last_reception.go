package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CloseLastReception(ctx context.Context, token token.Payload, pvzId uuid.UUID) (entity.Reception, error) {
	if token.UserRole != entity.RoleEmployee {
		return entity.Reception{}, entity.ErrUnauthorized
	}

	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockNoWrites)
	if err != nil {
		return entity.Reception{}, err
	}
	defer tx.Rollback()

	reception, err := u.repo.GetOpenReception(ctx, pvzId)
	// includes entity.ErrNotFound
	if err != nil {
		return entity.Reception{}, err
	}

	reception, err = u.repo.UpdateReceptionStatus(ctx, reception.Id, entity.StatusClose)
	if err != nil {
		return entity.Reception{}, err
	}

	return reception, tx.Commit()
}
