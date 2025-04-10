package repository

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/starnuik/avito_pvz/pkg/repository Repository
//go:generate mockgen -destination=../mocks/mock_repository_tx.go -package=mocks github.com/starnuik/avito_pvz/pkg/repository Tx

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

// TODO doc
type Repository interface {
	// Lock
	LockPvz(ctx context.Context, id uuid.UUID, lock DbLock) (Tx, error)

	// Create
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	CreatePvz(ctx context.Context, pvz entity.Pvz) (entity.Pvz, error)
	CreateReception(ctx context.Context, reception entity.Reception) (entity.Reception, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)

	// Read
	GetUser(ctx context.Context, email string) (entity.User, error)
	GetOpenReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error)
	GetLastProduct(ctx context.Context, pvzId uuid.UUID) (entity.Product, error)
	GetPvzInfo(ctx context.Context, startDate time.Time, endDate time.Time, limit int, offset int) (entity.PvzInfo, error)

	// Update
	UpdateReceptionStatus(ctx context.Context, id uuid.UUID, status entity.ReceptionStatus) (entity.Reception, error)

	// Delete
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

// TODO doc
type Tx interface {
	Commit() error
	Rollback() error
}

/*
Tx {
	LockPvz
		LastReception
		CreateReception
		UpdateReceptionStatus
	LockReception
		CreateProduct
		DeleteProduct
		LastProduct
}
*/

var _ Repository = (*pgImpl)(nil)

type pgImpl struct {
	conn *pgx.Conn
}

func (repo *pgImpl) GetPvzInfo(ctx context.Context, startDate time.Time, endDate time.Time, limit int, offset int) (entity.PvzInfo, error) {
	panic("unimplemented")
}

func (repo *pgImpl) GetLastProduct(ctx context.Context, pvzId uuid.UUID) (entity.Product, error) {
	panic("unimplemented")
}

func (repo *pgImpl) GetOpenReception(ctx context.Context, pvzId uuid.UUID) (entity.Reception, error) {
	panic("unimplemented")
}

// dont create multiple receptions
func (repo *pgImpl) LockPvz(ctx context.Context, id uuid.UUID, lock DbLock) (Tx, error) {
	panic("unimplemented")
}

func New(ctx context.Context, connString string) (*pgImpl, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &pgImpl{
		conn: conn,
	}, nil
}

type DbLock int

const (
	LockAllowWrites DbLock = iota
	LockNoWrites
)

/*
TODO
how do i treat optional json parameters?
ignore them?
send them to the db?
SOLUTION
do not create uuid's in the db, set them as 'not null' only
*/

// TODO
// func (repo *Repository) GetInfo()

// TODO
// helpers

// // required by UpdateReceptionStatus
// func (repo *pgImpl) LockPvz(ctx context.Context, id uuid.UUID) error {
// 	panic("")
// }

// // required by DeleteLastProduct
// func (repo *pgImpl) LockReception(ctx context.Context, id uuid.UUID) error {
// 	panic("")
// }

// // required by: CreateReception,
// func (repo *pgImpl) LastReception(ctx context.Context) (entity.Reception, error) {
// 	panic("")
// }

// func (repo *pgImpl) LastProduct(ctx context.Context, receptionId uuid.UUID) (entity.Product, error) {
// 	panic("")
// }
