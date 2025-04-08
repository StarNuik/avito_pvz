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

func Test_CreateProduct_ReceptionExists(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{
		Id:               pvztest.NewUuid(t),
		RegistrationDate: time.Now().UTC(),
		City:             entity.PvzCity(0),
	}
	testRepo.CreatePvz(t, pvz)

	reception := entity.Reception{
		Id:       uuid.Nil,
		PvzId:    pvz.Id,
		DateTime: time.Now().UTC(),
		Status:   entity.ReceptionStatus(0),
	}
	testRepo.CreateReception(t, reception)

	product := entity.Product{
		Id:          uuid.Nil,
		DateTime:    time.UnixMicro(0),
		ReceptionId: reception.Id,
		Type:        entity.ProductType(0),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.CreateProduct(ctx, product)

	// Assert
	require.Nil(err)

	require.NotEqual(uuid.Nil, result.Id)
	require.NotEqual(time.UnixMicro(0), result.DateTime)

	require.Equal(product.ReceptionId, result.ReceptionId)
	require.Equal(product.Type, result.Type)
}

func Test_CreateProduct_NoReception(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	product := entity.Product{
		Id:          uuid.Nil,
		DateTime:    time.UnixMicro(0),
		ReceptionId: uuid.Nil,
		Type:        entity.ProductType(0),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	_, err := repo.CreateProduct(ctx, product)

	// Assert
	require.Error(err)
}
