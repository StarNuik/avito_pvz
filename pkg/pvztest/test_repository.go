package pvztest

import (
	"context"
	"testing"

	"github.com/google/uuid"
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

func ClearRepository(t *testing.T) {
	repo := NewTestRepository(t)
	repo.Clear(t)
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

func (repo *TestRepository) CreateReception(t *testing.T, reception entity.Reception) {
	ctx := context.Background()
	_, err := repo.conn.Exec(ctx, `
		insert into receptions (id, pvzId, dateTime, status)
		values ($1, $2, $3, $4)
	`, reception.Id, reception.PvzId, reception.DateTime, reception.Status)

	if err != nil {
		t.Fatal(err)
	}
}

func (repo *TestRepository) CreateProduct(t *testing.T, product entity.Product) {
	ctx := context.Background()
	_, err := repo.conn.Exec(ctx, `
		insert into products (id, dateTime, receptionId, type)
		values ($1, $2, $3, $4)
	`, product.Id, product.DateTime, product.ReceptionId, product.Type)

	if err != nil {
		t.Fatal(err)
	}
}

func (repo *TestRepository) GetProduct(t *testing.T, productId uuid.UUID) entity.Product {
	ctx := context.Background()
	row := repo.conn.QueryRow(ctx, `
		select id, dateTime, receptionId, type
		from products
		where id = $1
	`, productId)

	product := entity.Product{}
	err := row.Scan(&product.Id, &product.DateTime, &product.ReceptionId, &product.Type)
	if err != nil {
		t.Fatal(err)
	}

	return product
}

func (repo *TestRepository) GetPvz(t *testing.T, pvzId uuid.UUID) entity.Pvz {
	ctx := context.Background()
	row := repo.conn.QueryRow(ctx, `
		select id, registrationDate, city
		from pvzs
		where id = $1
	`, pvzId)

	pvz := entity.Pvz{}
	err := row.Scan(&pvz.Id, &pvz.RegistrationDate, &pvz.City)
	if err != nil {
		t.Fatal(err)
	}

	return pvz
}

func (repo *TestRepository) GetReception(t *testing.T, receptionId uuid.UUID) entity.Reception {
	ctx := context.Background()
	row := repo.conn.QueryRow(ctx, `
		select id, pvzId, dateTime, status
		from receptions
		where id = $1
	`, receptionId)

	reception := entity.Reception{}
	err := row.Scan(&reception.Id, &reception.PvzId, &reception.DateTime, &reception.Status)
	if err != nil {
		t.Fatal(err)
	}

	return reception
}
