package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/gen"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

// TODO doc
type Usecase interface {
	// Auth
	Login(ctx context.Context, email string, password string) (token.Payload, error)
	Register(ctx context.Context, email string, password string, role entity.UserRole) (entity.User, error)
	DummyLogin(userRole entity.UserRole) token.Payload

	// Create
	CreatePvz(ctx context.Context, token token.Payload, city entity.PvzCity, id *uuid.UUID, registrationDate *time.Time) (entity.Pvz, error)
	CreateProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error)
	CreateReception(tctx context.Context, token token.Payload, pvzId uuid.UUID) (entity.Reception, error)

	// Read
	GetPvz(ctx context.Context, token token.Payload, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error)

	// Update
	CloseLastReception(ctx context.Context, token token.Payload, pvzId uuid.UUID) (entity.Reception, error)

	// Delete
	DeleteLastProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID) error
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	repo   repository.Repository
	hasher password.Hasher
	gen    gen.Gen
}

func New(repo repository.Repository, hasher password.Hasher, gen gen.Gen) *usecase {
	return &usecase{
		repo:   repo,
		hasher: hasher,
		gen:    gen,
	}
}
