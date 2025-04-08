package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *Repository) CreateReception(ctx context.Context, reception entity.Reception) (entity.Reception, error) {
	row := repo.conn.QueryRow(ctx, `
		insert into receptions (pvzId, dateTime, status)
		values ($1, $2, $3)
		returning id, dateTime
	`, reception.PvzId, reception.DateTime, reception.Status)

	err := row.Scan(&reception.Id, &reception.DateTime)
	return reception, err
}
