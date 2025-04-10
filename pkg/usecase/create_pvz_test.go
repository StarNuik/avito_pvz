package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_CreatePvz_AllOptions(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	pvz := entity.Pvz{
		Id:               uuid.Max,
		City:             entity.PvzCity(1),
		RegistrationDate: time.Unix(1000, 0),
	}

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		CreatePvz(gomock.Any(), pvz).
		Return(nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleModerator,
	}

	// Act
	result, err := usecase.CreatePvz(ctx, token, pvz.City, &pvz.Id, &pvz.RegistrationDate)

	// Assert
	require.Nil(err)
	require.Equal(pvz, result)
}

func Test_CreatePvz_Unauthorized(t *testing.T) {
	// Arrange
	require := require.New(t)

	usecase := usecase.New(nil, nil, nil)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleEmployee,
	}

	// Act
	_, err := usecase.CreatePvz(ctx, token, entity.PvzCity(1), nil, nil)

	// Assert
	require.ErrorIs(err, entity.ErrUnauthorized)
}
