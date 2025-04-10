package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

type DbLock int

const (
	LockAllowWrites DbLock = iota
	LockNoWrites
)

func (repo *pgImpl) LockPvz(ctx context.Context, pvzId uuid.UUID, lock DbLock) (Tx, error) {
	/*
		TODO remove me

		# usecases
		create reception VS create reception
		create product VS close reception
		create product VS delete product
		delete product VS delete product
	*/

	tx, err := repo.conn.Begin(ctx)
	if err != nil {
		return nil, entity.InternalError("LockPvz", err)
	}

	sqlLock := lockToSql(lock)
	sql := fmt.Sprintf(`
		select from pvzs
		where id = $1
		%s`, sqlLock)

	row := repo.conn.QueryRow(ctx, sql, pvzId)
	err = row.Scan()
	if errors.Is(err, pgx.ErrNoRows) {
		tx.Rollback(ctx)
		return nil, entity.ErrNotFound
	}
	if err != nil {
		tx.Rollback(ctx)
		return nil, entity.InternalError("LockPvz", err)
	}

	return tx, err
}

func lockToSql(lock DbLock) string {
	switch lock {
	case LockAllowWrites:
		return "for key share"
	case LockNoWrites:
		return "for update"
	default:
		panic("lock not implemented")
	}
}
