package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreateProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error) {
	if token.UserRole != entity.RoleEmployee {
		return entity.Product{}, entity.ErrUnauthorized
	}

	tx, err := u.repo.LockPvz(ctx, pvzId, repository.LockAllowWrites)
	if err != nil {
		return entity.Product{}, nil
	}
	defer tx.Rollback()

	reception, err := u.repo.GetOpenReception(ctx, pvzId)
	if err != nil {
		return entity.Product{}, err
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

	product, err = u.repo.CreateProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, tx.Commit()
}
