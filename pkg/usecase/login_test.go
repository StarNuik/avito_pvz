package usecase_test

import (
	"context"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_Login_HappyPath(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	user := entity.User{
		Email:        "email",
		PasswordHash: []byte("hashed password"),
		Role:         entity.UserRole(1),
	}
	password := "password"

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(user.Email)).
		Return(user, nil)

	hasher := mocks.NewMockHasher(ctrl)
	hasher.EXPECT().
		Compare(gomock.Eq(password), gomock.Eq(user.PasswordHash)).
		Return(true)

	ctx := context.Background()
	usecase := usecase.New(repo, hasher)

	// Act
	result, err := usecase.Login(ctx, user.Email, password)

	// Assert
	require.Nil(err)
	require.Equal(user.Role, result.UserRole)
}

func Test_Login_UserNotFound(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		GetUser(gomock.Any(), gomock.Any()).
		Return(entity.User{}, entity.ErrNotFound)

	ctx := context.Background()
	usecase := usecase.New(repo, nil)

	// Act
	_, err := usecase.Login(ctx, "email", "password")

	// Assert
	require.ErrorIs(err, entity.ErrIncorrectLogin)
}

func Test_Login_PasswordDoesntMatch(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	user := entity.User{
		Email:        "email",
		PasswordHash: []byte("hashed password"),
		Role:         entity.UserRole(1),
	}

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(user.Email)).
		Return(user, nil)

	hasher := mocks.NewMockHasher(ctrl)
	hasher.EXPECT().
		Compare(gomock.Any(), gomock.Any()).
		Return(false)

	ctx := context.Background()
	usecase := usecase.New(repo, hasher)

	// Act
	_, err := usecase.Login(ctx, user.Email, "password")

	// Assert
	require.ErrorIs(err, entity.ErrIncorrectLogin)
}
