package usecase

import (
	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/repository"
)

type Usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// func (u *Usecase) GetPvz() error {}

func (u *Usecase) CreatePvz(user entity.UserRole, pvz entity.Pvz) (entity.Pvz, error) {
	panic("")
}

func (u *Usecase) CreateReception(user entity.UserRole, reception entity.Reception) (entity.Reception, error) {
	panic("")
}

func (u *Usecase) CreateProduct(user entity.UserRole, product entity.Product) (entity.Product, error) {
	panic("")
}

func (u *Usecase) CloseLastReception(user entity.UserRole, pvzId uuid.UUID) (entity.Reception, error) {
	panic("")
}

func (u *Usecase) DeleteLastProduct(user entity.UserRole, pvzId uuid.UUID) error {
	panic("")
}
