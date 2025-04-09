package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CreateProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error) {
	if token.UserRole != entity.RoleEmployee {
		return entity.Product{}, entity.ErrUnauthorized
	}

	panic("")
	// id, err := u.gen.Uuid()
	// if err != nil {
	// 	return entity.Product{}, err
	// }

	// now := u.gen.Now()

	// receptionId := 0

	// product := entity.Product{
	// 	Id:          id,
	// 	DateTime:    now,
	// 	ReceptionId: receptionId,
	// 	Type:        productType,
	// }
}
