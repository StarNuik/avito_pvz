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

func Test_CreateReception(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	reception := entity.Reception{
		Id:       uuid.Max,
		PvzId:    uuid.Max,
		DateTime: time.Unix(1000, 0),
		Status:   entity.StatusInProgress,
	}

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().Rollback().After(tx.EXPECT().Commit())

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), reception.PvzId).
		Return(tx, nil)
	repo.EXPECT().
		GetOpenReception(gomock.Any(), reception.PvzId).
		Return(entity.Reception{}, entity.ErrNotFound)
	repo.EXPECT().
		CreateReception(gomock.Any(), reception).
		Return(reception, nil)

	gen := mocks.NewMockGen(ctrl)
	gen.EXPECT().
		Now().
		Return(reception.DateTime)
	gen.EXPECT().
		Uuid().
		Return(reception.Id, nil)

	usecase := usecase.New(repo, nil, gen)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleEmployee,
	}

	// Act
	result, err := usecase.CreateReception(ctx, token, reception.PvzId)

	// Assert
	require.Nil(err)
	require.Equal(reception, result)
}

func Test_CreateReception_Unauthorized(t *testing.T) {
	// Arrange
	require := require.New(t)

	usecase := usecase.New(nil, nil, nil)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleModerator,
	}

	// Act
	_, err := usecase.CreateReception(ctx, token, uuid.Nil)

	// Assert
	require.ErrorIs(err, entity.ErrUnauthorized)
}
