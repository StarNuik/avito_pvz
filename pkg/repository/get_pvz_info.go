package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) GetPvzInfo(ctx context.Context, startDate time.Time, endDate time.Time, limit int, offset int) (entity.PvzInfo, error) {
	rows, err := repo.conn.Query(ctx, `
		select
			pvz.id, pvz.registrationDate, pvz.city,
			r.id, r.pvzId, r.dateTime, r.status,
			p.id, p.dateTime, p.receptionId, p.type
		from products p
		join receptions r
			on p.receptionId = r.id
		join pvzs pvz
			on r.pvzId = pvz.id
		where p.dateTime > $1 and p.dateTime < $2
		limit $3
		offset $4
	`, startDate, endDate, limit, offset)
	if err != nil {
		return entity.PvzInfo{}, entity.InternalError("GetPvzInfo", err)
	}

	out := entity.PvzInfo{
		Pvzs:       make(map[uuid.UUID]entity.Pvz),
		Receptions: make(map[uuid.UUID]entity.Reception),
		Products:   make(map[uuid.UUID]entity.Product),
	}
	for rows.Next() {
		pvz := entity.Pvz{}
		reception := entity.Reception{}
		product := entity.Product{}

		err := rows.Scan(&pvz.Id, &pvz.RegistrationDate, &pvz.City,
			&reception.Id, &reception.PvzId, &reception.DateTime, &reception.Status,
			&product.Id, &product.DateTime, &product.ReceptionId, &product.Type)
		if err != nil {
			return entity.PvzInfo{}, entity.InternalError("GetPvzInfo", err)
		}

		out.Pvzs[pvz.Id] = pvz
		out.Receptions[reception.Id] = reception
		out.Products[product.Id] = product
	}

	return out, nil
}
