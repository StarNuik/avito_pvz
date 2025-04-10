package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_CloseLastReception(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	reception := entity.Reception{
		Id:       uuid.Max,
		PvzId:    uuid.Max,
		Status:   entity.StatusClose,
		DateTime: time.Unix(1000, 0),
	}

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().Rollback().After(tx.EXPECT().Commit())

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), reception.PvzId, repository.LockNoWrites).
		Return(tx, nil)
	repo.EXPECT().
		GetLastReception(gomock.Any(), reception.PvzId).
		Return(reception, nil)
	repo.EXPECT().
		UpdateReceptionStatus(gomock.Any(), reception.Id, entity.StatusClose).
		Return(reception, nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleEmployee,
	}

	// Act
	result, err := usecase.CloseLastReception(ctx, token, reception.PvzId)

	// Assert
	require.Nil(err)
	require.Equal(reception, result)
}
