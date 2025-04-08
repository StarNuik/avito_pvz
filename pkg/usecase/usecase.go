package usecase

import "github.com/starnuik/avito_pvz/pkg/repository"

type Usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// func (u *Usecase) CreatePvz()
