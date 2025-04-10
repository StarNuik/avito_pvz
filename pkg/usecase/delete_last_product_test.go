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

	pvzId := uuid.Max
	product := entity.Product{
		Id:          uuid.Max,
		ReceptionId: uuid.Max,
		DateTime:    time.Unix(1000, 0),
		Type:        entity.ProductType(128),
	}

	tx := mocks.NewMockTx(ctrl)
	tx.EXPECT().Rollback().After(tx.EXPECT().Commit())

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		LockPvz(gomock.Any(), pvzId, repository.LockNoWrites).
		Return(tx, nil)
	repo.EXPECT().
		GetLastProduct(gomock.Any(), pvzId).
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
	err := usecase.DeleteLastProduct(ctx, token, pvzId)

	// Assert
	require.Nil(err)
}
