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

func Test_CreateProduct(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	reception := entity.Reception{
		Id:    uuid.Max,
		PvzId: uuid.Max,
	}
	product := entity.Product{
		Id:          uuid.Max,
		ReceptionId: reception.Id,
		DateTime:    time.Unix(1000, 0),
		Type:        entity.ProductType(1),
	}

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().
		Rollback(gomock.Any()).
		After(tx.EXPECT().
			Commit(gomock.Any()))

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), reception.PvzId, repository.LockAllowWrites).
		Return(tx, nil)
	repo.EXPECT().
		GetLastReception(gomock.Any(), reception.PvzId).
		Return(reception, nil)
	repo.EXPECT().
		CreateProduct(gomock.Any(), product).
		Return(nil)

	gen := mocks.NewMockGen(ctrl)
	gen.EXPECT().
		Now().
		Return(product.DateTime)
	gen.EXPECT().
		Uuid().
		Return(product.Id, nil)

	usecase := usecase.New(repo, nil, gen)

	ctx := context.Background()

	// Act
	result, err := usecase.CreateProduct(ctx, reception.PvzId, product.Type)

	// Assert
	require.Nil(err)
	require.Equal(product, result)
}
