package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_Register(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	password := "password"
	user := entity.User{
		Id:           uuid.Max,
		Email:        "new@email.com",
		Role:         entity.UserRole(1),
		PasswordHash: []byte("bytes"),
	}

	gen := mocks.NewMockGen(ctrl)
	gen.EXPECT().
		Uuid().
		Return(user.Id, nil)

	hasher := mocks.NewMockHasher(ctrl)
	hasher.EXPECT().
		Hash(password).
		Return(user.PasswordHash, nil)

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().CreateUser(gomock.Any(), user).
		Return(nil)

	ctx := context.Background()

	usecase := usecase.New(repo, hasher, gen)

	// Act
	result, err := usecase.Register(ctx, user.Email, password, entity.UserRole(1))

	// Assert
	require.Nil(err)
	require.Equal(user, result)
}
