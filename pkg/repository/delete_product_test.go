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

func Test_DeleteProduct_ProductExists(t *testing.T) {
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
		Id:          pvztest.NewUuid(t),
		DateTime:    time.Now().UTC(),
		ReceptionId: reception.Id,
		Type:        entity.ProductType(0),
	}
	testRepo.CreateProduct(t, product)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.DeleteProduct(ctx, product.Id)

	// Assert
	require.Nil(err)
}

func Test_DeleteProduct_NoProduct(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.DeleteProduct(ctx, uuid.Nil)

	// Assert
	require.Error(err)
}
