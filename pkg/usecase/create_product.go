package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreateProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error) {
	panic("")
}
