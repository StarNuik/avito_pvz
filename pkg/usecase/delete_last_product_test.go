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

func Test_DeleteLastProduct(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	reception := entity.Reception{}
	product := entity.Product{
		Id:          uuid.Max,
		ReceptionId: reception.Id,
		DateTime:    time.Unix(1000, 0),
		Type:        entity.ProductType(128),
	}

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().Rollback().After(tx.EXPECT().Commit())

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), reception.PvzId, repository.LockNoWrites).
		Return(tx, nil)
	repo.EXPECT().GetLastReception(gomock.Any(), reception.PvzId).
		Return(reception, nil)
	repo.EXPECT().
		GetLastProduct(gomock.Any(), reception.Id).
		Return(product, nil)
	repo.EXPECT().
		DeleteProduct(gomock.Any(), product.Id).
		Return(nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()
	token := token.Payload{
		UserRole: entity.RoleEmployee,
	}

	// Act
	err := usecase.DeleteLastProduct(ctx, token, reception.Id)

	// Assert
	require.Nil(err)
}
