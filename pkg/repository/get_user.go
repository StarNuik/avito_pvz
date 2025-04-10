package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (repo *pgImpl) GetUser(ctx context.Context, email string) (entity.User, error) {
	row := repo.conn.QueryRow(ctx, `
		select id, email, role, passwordHash
		from users
		where email = $1
	`, email)

	user := entity.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Role, &user.PasswordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.User{}, entity.ErrNotFound
	}
	if err != nil {
		return entity.User{}, entity.InternalError("GetUser", err)
	}

	return user, nil
}
