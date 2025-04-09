package usecase

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (u *usecase) Register(ctx context.Context, email string, password string, role entity.UserRole) (entity.User, error) {
	panic("")
}
