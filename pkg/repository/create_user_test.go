package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_CreateUser(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	user := entity.User{
		Id:           uuid.Nil,
		Email:        "test@email.com",
		PasswordHash: []byte("weak password"),
		Role:         entity.UserRole(0),
	}
	repo := pvztest.NewRepository(t)
	ctx := context.Background()

	// Act
	err := repo.CreateUser(ctx, user)

	// Assert
	require.Nil(err)
}
