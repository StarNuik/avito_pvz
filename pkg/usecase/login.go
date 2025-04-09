package usecase

import (
	"context"
	"errors"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) Login(ctx context.Context, email string, password string) (token.Payload, error) {

	user, err := u.repo.GetUser(ctx, email)
	if errors.Is(err, entity.ErrNotFound) {
		return token.Payload{}, entity.ErrIncorrectLogin
	}
	if err != nil {
		return token.Payload{}, err
	}

	passOk := u.hasher.Compare(password, user.PasswordHash)
	if !passOk {
		return token.Payload{}, entity.ErrIncorrectLogin
	}

	return token.Payload{
		UserRole: user.Role,
	}, nil
}
