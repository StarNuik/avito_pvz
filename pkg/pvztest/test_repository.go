package pvztest

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

type TestRepository struct {
	conn *pgx.Conn
}

const testConnString = "postgres://postgres:postgres@localhost:5432/pvz"

func NewTestRepository(t *testing.T) *TestRepository {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, testConnString)
	if err != nil {
		t.Fatal(err)
	}
	return &TestRepository{
		conn: conn,
	}
}

// func (repo *TestRepository) GetUser(t *testing.T, id uuid.UUID) entity.User {
// 	ctx := context.Background()
// 	row := repo.conn.QueryRow(ctx, `
// 		select id, email, role, passwordHash
// 		from users
// 		where id = $1
// 	`, id)

// 	user := entity.User{}
// 	err := row.Scan(&user.Id, &user.Email, &user.Role, &user.PasswordHash)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	return user
// }

func (repo *TestRepository) Clear(t *testing.T) {
	ctx := context.Background()
	_, err := repo.conn.Exec(ctx, `
		truncate table users, pvzs, receptions, products;
	`)
	if err != nil {
		t.Fatal(err)
	}
}

func (repo *TestRepository) CreateUser(t *testing.T, user entity.User) {
	ctx := context.Background()
	_, err := repo.conn.Exec(ctx, `
		insert into users (id, email, role, passwordHash)
		values ($1, $2, $3, $4)
	`, user.Id, user.Email, user.Role, user.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}
}

func (repo *TestRepository) CreatePvz(t *testing.T, pvz entity.Pvz) {
	ctx := context.Background()
	_, err := repo.conn.Exec(ctx, `
		insert into pvzs (id, registrationDate, city)
		values ($1, $2, $3)
	`, pvz.Id, pvz.RegistrationDate, pvz.City)

	if err != nil {
		t.Fatal(err)
	}
}
