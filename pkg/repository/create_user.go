package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreateUser(ctx context.Context, user entity.User) error {
	_, err := repo.conn.Exec(ctx, `
		insert into users (id, email, role, passwordHash)
		values ($1, $2, $3, $4)
	`, user.Id, user.Email, user.Role, user.PasswordHash)
	if err != nil {
		return entity.InternalError("CreateUser", err)
	}

	return nil
}
