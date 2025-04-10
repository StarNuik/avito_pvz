package usecase

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (u *usecase) Register(ctx context.Context, email string, password string, role entity.UserRole) (entity.User, error) {
	userId, err := u.gen.Uuid()
	if err != nil {
		return entity.User{}, err
	}

	passwordHash, err := u.hasher.Hash(password)
	if err != nil {
		return entity.User{}, nil
	}

	user := entity.User{
		Id:           userId,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}
	err = u.repo.CreateUser(ctx, user)

	return user, err
}
