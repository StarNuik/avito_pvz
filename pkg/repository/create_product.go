package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreateProduct(ctx context.Context, product entity.Product) error {
	_, err := repo.conn.Exec(ctx, `
		insert into products (id, dateTime, receptionId, type)
		values ($1, $2, $3, $4)
	`, product.Id, product.DateTime, product.ReceptionId, product.Type)
	if err != nil {
		return entity.InternalError("CreateProduct", err)
	}

	return nil
}
