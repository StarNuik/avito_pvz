package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreatePvz(ctx context.Context, pvz entity.Pvz) error {
	_, err := repo.conn.Exec(ctx, `
		insert into pvzs (id, registrationDate, city)
		values ($1, $2, $3)
	`, pvz.Id, pvz.RegistrationDate, pvz.City)
	if err != nil {
		return entity.InternalError("CreatePvz", err)
	}

	return nil
}
