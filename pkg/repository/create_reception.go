package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreateReception(ctx context.Context, reception entity.Reception) error {
	_, err := repo.conn.Exec(ctx, `
		insert into receptions (id, pvzId, dateTime, status)
		values ($1, $2, $3, $4)
	`, reception.Id, reception.PvzId, reception.DateTime, reception.Status)
	if err != nil {
		return entity.InternalError("CreateReception", err)
	}

	return nil
}
