package repository

import (
	"context"

	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *Repository) GetUser(ctx context.Context, email string) (entity.User, error) {
	row := repo.conn.QueryRow(ctx, `
		select id, email, role, passwordHash
		from users
		where email = $1
	`, email)

	user := entity.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Role, &user.PasswordHash)
	return user, err
}
