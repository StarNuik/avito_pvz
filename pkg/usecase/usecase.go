package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
)

// TODO doc
type Usecase interface {
	// Auth
	Login(email string, password string) (token.Payload, error)
	Register(email string, password string, role entity.UserRole) (entity.User, error)
	DummyLogin(userRole entity.UserRole) (token.Payload, error)

	// Create
	CreatePvz(token token.Payload, city entity.PvzCity, id *uuid.UUID, registrationDate *time.Time) (entity.Pvz, error)
	CreateProduct(token token.Payload, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error)
	CreateReception(token token.Payload, pvzId uuid.UUID) (entity.Reception, error)

	// Read
	GetPvz(token token.Payload, startDate time.Time, endDate time.Time, page *int, limit *int) (entity.PvzInfo, error)

	// Update
	CloseLastReception(token token.Payload, pvzId uuid.UUID) (entity.Reception, error)

	// Delete
	DeleteLastProduct(token token.Payload, pvzId uuid.UUID) error
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}
