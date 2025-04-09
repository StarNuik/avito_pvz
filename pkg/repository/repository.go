package repository

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/starnuik/avito_pvz/pkg/repository Repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

// TODO doc
type Repository interface {
	// Create
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	CreatePvz(ctx context.Context, pvz entity.Pvz) (entity.Pvz, error)
	CreateReception(ctx context.Context, reception entity.Reception) (entity.Reception, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)

	// Read
	GetUser(ctx context.Context, email string) (entity.User, error)

	// Update
	UpdateReceptionStatus(ctx context.Context, id uuid.UUID, status entity.ReceptionStatus) (entity.Reception, error)

	// Delete
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

var _ Repository = (*pgImpl)(nil)

type pgImpl struct {
	conn *pgx.Conn
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

// TODO
// how do i treat optional json parameters?
// ignore them?
// send them to the db?
// SOLUTION
// do not create uuid's in the db, set them as 'not null' only

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

// func (repo *pgImpl) LastReception(ctx context.Context) (entity.Reception, error) {
// 	panic("")
// }

// func (repo *pgImpl) LastProduct(ctx context.Context, receptionId uuid.UUID) (entity.Product, error) {
// 	panic("")
// }
