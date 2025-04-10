package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) GetLastReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error) {
	row := repo.conn.QueryRow(ctx, `
		select id, pvzId, dateTime, status
		from receptions
		where pvzId = $1
		order by dateTime desc
		limit 1
	`, pvzId)

	reception := entity.Reception{}
	err := row.Scan(&reception.Id, &reception.PvzId, &reception.DateTime, &reception.Status)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Reception{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.Reception{}, entity.InternalError("GetLastReception", err)
	}

	return reception, nil
}
