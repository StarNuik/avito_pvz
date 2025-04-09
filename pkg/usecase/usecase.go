package usecase

import (
	"go/token"
	"time"

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

// Auth
func (u *Usecase) DummyLogin(userRole entity.UserRole) token.Token {
	panic("")
}

func (u *Usecase) Register(email string, password string, role entity.UserRole) entity.User {
	panic("")
}

func (u *Usecase) Login(email string, password string) token.Token {
	panic("")
}

// Create
func (u *Usecase) CreatePvz(reqRole entity.UserRole, city entity.PvzCity, id *uuid.UUID, registrationDate *time.Time) (entity.Pvz, error) {
	panic("")
}

func (u *Usecase) CreateReception(reqRole entity.UserRole, pvzId uuid.UUID) (entity.Reception, error) {
	panic("")
}

func (u *Usecase) CreateProduct(reqRole entity.UserRole, pvzId uuid.UUID, productType entity.ProductType) (entity.Product, error) {
	panic("")
}

// Read
// func (u *Usecase) GetPvz() error {}

// Update / Delete
func (u *Usecase) CloseLastReception(reqRole entity.UserRole, pvzId uuid.UUID) (entity.Reception, error) {
	panic("")
}

func (u *Usecase) DeleteLastProduct(reqRole entity.UserRole, pvzId uuid.UUID) error {
	panic("")
}
