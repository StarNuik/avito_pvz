package repository

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/starnuik/avito_pvz/pkg/repository Repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

// TODO doc
type Repository interface {
	Close(context.Context) error

	// Lock
	LockPvz(ctx context.Context, id uuid.UUID, lock DbLock) (Tx, error)

	// Create
	CreateProduct(ctx context.Context, product entity.Product) error
	CreatePvz(ctx context.Context, pvz entity.Pvz) error
	CreateReception(ctx context.Context, reception entity.Reception) error
	CreateUser(ctx context.Context, user entity.User) error

	// Read
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetLastReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error)
	GetLastProduct(ctx context.Context, receptionId uuid.UUID) (entity.Product, error)
	GetPvzInfo(ctx context.Context, startDate time.Time, endDate time.Time, limit int, offset int) (entity.PvzInfo, error)

	// Update
	UpdateReceptionStatus(ctx context.Context, id uuid.UUID, status entity.ReceptionStatus) (entity.Reception, error)

	// Delete
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

var _ Repository = (*pgImpl)(nil)

type pgImpl struct {
	conn  *pgx.Conn
	close func(context.Context) error
}

func New(ctx context.Context, connString string) (*pgImpl, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &pgImpl{
		conn:  conn,
		close: conn.Close,
	}, nil
}

func (repo *pgImpl) Close(ctx context.Context) error {
	return repo.close(ctx)
}
