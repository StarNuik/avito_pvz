package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	row := repo.conn.QueryRow(ctx, `
		insert into users (email, role, passwordHash)
		values ($1, $2, $3)
		returning id
	`, user.Email, user.Role, user.PasswordHash)

	err := row.Scan(&user.Id)
	return user, err
}
