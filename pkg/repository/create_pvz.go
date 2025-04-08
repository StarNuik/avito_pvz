package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreatePvz(ctx context.Context, pvz entity.Pvz) (entity.Pvz, error) {
	row := repo.conn.QueryRow(ctx, `
		insert into pvzs (registrationDate, city)
		values ($1, $2)
		returning id, registrationDate
	`, pvz.RegistrationDate, pvz.City)

	err := row.Scan(&pvz.Id, &pvz.RegistrationDate)
	return pvz, err
}
