package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_GetUser(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	userId, err := uuid.NewRandom()
	require.Nil(err)

	user := entity.User{
		Id:           userId,
		Email:        "test@email.com",
		PasswordHash: []byte("weak password"),
		Role:         entity.UserRole(0),
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
