package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_CreateProduct_ReceptionExists(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{
		Id: pvztest.NewUuid(t),
	}
	testRepo.CreatePvz(t, pvz)

	reception := entity.Reception{
		Id:    pvztest.NewUuid(t),
		PvzId: pvz.Id,
	}
	testRepo.CreateReception(t, reception)

	product := entity.Product{
		Id:          pvztest.NewUuid(t),
		DateTime:    time.Unix(1000, 0),
		ReceptionId: reception.Id,
		Type:        entity.ProductType(128),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.CreateProduct(ctx, product)

	// Assert
	require.Nil(err)

	result := testRepo.GetProduct(t, product.Id)
	require.Equal(product, result)
}

func Test_CreateProduct_NoReception(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	product := entity.Product{}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.CreateProduct(ctx, product)

	// Assert
	require.ErrorIs(err, entity.ErrInternal)
}
