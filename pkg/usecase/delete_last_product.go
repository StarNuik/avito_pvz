package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) DeleteLastProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID) error {
	if token.UserRole != entity.RoleEmployee {
		return entity.ErrUnauthorized
	}

	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockNoWrites)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	product, err := u.repo.GetLastProduct(ctx, pvzId)
	if err != nil {
		return err
	}

	err = u.repo.DeleteProduct(ctx, product.Id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
