package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) GetLastProduct(ctx context.Context, receptionId uuid.UUID) (entity.Product, error) {
	row := repo.conn.QueryRow(ctx, `
		select id, dateTime, receptionId, type
		from products
		where receptionId = $1
		order by dateTime desc
		limit 1
	`, receptionId)

	product := entity.Product{}
	err := row.Scan(&product.Id, &product.DateTime, &product.ReceptionId, &product.Type)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Product{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.Product{}, entity.InternalError("GetLastProduct", err)
	}

	return product, nil
}
