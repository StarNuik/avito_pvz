package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *Repository) UpdateReceptionStatus(ctx context.Context, id uuid.UUID, status entity.ReceptionStatus) (entity.Reception, error) {
	row := repo.conn.QueryRow(ctx, `
		update receptions
		set status = $1
		where id = $2
		returning id, pvzId, dateTime, status
	`, status, id)

	reception := entity.Reception{}
	err := row.Scan(&reception.Id, &reception.PvzId, &reception.DateTime, &reception.Status)
	return reception, err
}
