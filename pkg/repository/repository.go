package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(ctx context.Context, connString string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &Repository{
		conn: conn,
	}, nil
}

// TODO
// func (repo *Repository) GetInfo()
// helpers
func (repo *Repository) LockPvz(ctx context.Context, id uuid.UUID) error {
	panic("")
}

func (repo *Repository) LockReception(ctx context.Context, id uuid.UUID) error {
	panic("")
}

func (repo *Repository) LastReception(ctx context.Context) (entity.Reception, error) {
	panic("")
}

func (repo *Repository) LastProduct(ctx context.Context, receptionId uuid.UUID) (entity.Product, error) {
	panic("")
}
