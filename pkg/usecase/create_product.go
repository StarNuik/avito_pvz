package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
)

func (u *usecase) CreateProduct(ctx context.Context, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error) {
	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockAllowWrites)
	if err != nil {
		return entity.Product{}, nil
	}
	defer tx.Rollback(ctx)

	reception, err := u.repo.GetLastReception(ctx, pvzId)
	if err != nil {
		return entity.Product{}, err
	}

	if reception.Status != entity.StatusInProgress {
		return entity.Product{}, entity.ErrReceptionClosed
	}

	id, err := u.gen.Uuid()
	if err != nil {
		return entity.Product{}, err
	}

	now := u.gen.Now()

	product := entity.Product{
		Id:          id,
		DateTime:    now,
		ReceptionId: reception.Id,
		Type:        productType,
	}

	err = u.repo.CreateProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, tx.Commit(ctx)
}
