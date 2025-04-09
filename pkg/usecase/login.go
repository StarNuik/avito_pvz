package usecase

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) Login(ctx context.Context, email string, password string) (token.Payload, error) {
	out := token.Payload{}

	_, err := u.repo.GetUser(ctx, email)
	if err != nil {
		return out, err
	}

	panic("")
}
