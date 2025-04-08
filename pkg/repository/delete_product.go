package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (repo *Repository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	tag, err := repo.conn.Exec(ctx, `
		delete from products
		where id = $1
	`, id)

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}
