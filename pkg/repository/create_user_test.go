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

	repo := pvztest.NewRepository(t)

	ctx := context.Background()

	user := entity.User{
		Email:        "test@email.com",
		PasswordHash: []byte("weak hash"),
		Role:         entity.UserRole(0),
	}

	// Act
	result, err := repo.CreateUser(ctx, user)

	// Assert
	require.Nil(err)
	require.NotEqual(uuid.Nil, result.Id)
	require.Equal(user.Email, result.Email)
	require.Equal(user.Role, result.Role)
	require.Equal(user.PasswordHash, result.PasswordHash)

	testRepo := pvztest.NewTestRepository(t)
	dbUser := testRepo.GetUser(t, result.Id)

	require.Equal(result, dbUser)
}
