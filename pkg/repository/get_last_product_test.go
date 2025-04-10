package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_GetLastProduct(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{}
	testRepo.CreatePvz(t, pvz)

	reception := entity.Reception{PvzId: pvz.Id}
	testRepo.CreateReception(t, reception)

	products := []entity.Product{
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(100, 0), ReceptionId: reception.Id},
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(200, 0), ReceptionId: reception.Id},
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(300, 0), ReceptionId: reception.Id},
	}
	for _, p := range products {
		testRepo.CreateProduct(t, p)
	}

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.GetLastProduct(ctx, reception.Id)

	// Assert
	require.Nil(err)
	require.Equal(products[2], result)
}

func Test_GetLastProduct_NotFound(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	_, err := repo.GetLastProduct(ctx, uuid.Nil)

	// Assert
	require.ErrorIs(err, entity.ErrNotFound)
}
