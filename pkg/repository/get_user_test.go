package repository_test

import (
	"context"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_GetUser(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	user := entity.User{
		Id:           pvztest.NewUuid(t),
		Email:        "test@email.com",
		PasswordHash: []byte("weak password"),
		Role:         entity.UserRole(128),
	}
	testRepo.CreateUser(t, user)

	repo := pvztest.NewRepository(t)
	ctx := context.Background()

	// Act
	result, err := repo.GetUser(ctx, user.Email)

	// Assert
	require.Nil(err)
	require.Equal(user, result)
}

func Test_GetUser_NotFound(t *testing.T) {
	// Arrange
	require := require.New(t)

	repo := pvztest.NewRepository(t)
	ctx := context.Background()

	// Act
	_, err := repo.GetUser(ctx, "not@found")

	// Assert
	require.ErrorIs(err, entity.ErrNotFound)
}
