package repo

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

func (repo *Repository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	panic("")
}

func (repo *Repository) GetUser(ctx context.Context, email string) (entity.User, error) {
	panic("")
}

func (repo *Repository) CreatePvz(ctx context.Context, pvz entity.Pvz) (entity.Pvz, error) {
	panic("")
}

// TODO
// func (repo *Repository) GetInfo()

func (repo *Repository) CloseReception(ctx context.Context, id uuid.UUID) error {
	panic("")
}

func (repo *Repository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	panic("")
}

func (repo *Repository) CreateReception(ctx context.Context, reception entity.Reception) (entity.Reception, error) {
	panic("")
}

func (repo *Repository) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	panic("")
}

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

func (repo *Repository) LastProduct(ctx context.Context) (entity.Product, error) {
	panic("")
}
