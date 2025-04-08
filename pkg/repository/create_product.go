package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *Repository) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	row := repo.conn.QueryRow(ctx, `
		insert into products (receptionId, type)
		values ($1, $2)
		returning id, dateTime, receptionId, type
	`, product.ReceptionId, product.Type)

	err := row.Scan(&product.Id, &product.DateTime, &product.ReceptionId, &product.Type)
	return product, err
}
