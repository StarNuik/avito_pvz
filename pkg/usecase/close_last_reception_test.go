package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_CloseLastReception(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	receptionOpen := entity.Reception{
		Id:       uuid.Max,
		PvzId:    uuid.Max,
		Status:   entity.StatusInProgress,
		DateTime: time.Unix(1000, 0),
	}
	receptionClosed := receptionOpen
	receptionClosed.Status = entity.StatusClosed

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().
		Rollback(gomock.Any()).
		After(tx.EXPECT().
			Commit(gomock.Any()))

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), receptionOpen.PvzId, repository.LockNoWrites).
		Return(tx, nil)
	repo.EXPECT().
		GetLastReception(gomock.Any(), receptionOpen.PvzId).
		Return(receptionOpen, nil)
	repo.EXPECT().
		UpdateReceptionStatus(gomock.Any(), receptionOpen.Id, entity.StatusClosed).
		Return(receptionClosed, nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()

	// Act
	result, err := usecase.CloseLastReception(ctx, receptionOpen.PvzId)

	// Assert
	require.Nil(err)
	require.Equal(receptionClosed, result)
}
