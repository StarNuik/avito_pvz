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

func Test_CreateReception_PvzExists(t *testing.T) {
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
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.CreateReception(ctx, reception)

	// Assert
	require.Nil(err)
	require.NotEqual(uuid.Nil, result.Id)

	require.Equal(reception.PvzId, result.PvzId)
	require.Equal(reception.Status, result.Status)
	pvztest.RequireEqualTime(t, reception.DateTime, result.DateTime)
}

func Test_CreateReception_NoPvz(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	reception := entity.Reception{
		Id:       uuid.Nil,
		PvzId:    uuid.Nil,
		DateTime: time.Now().UTC(),
		Status:   entity.ReceptionStatus(0),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	_, err := repo.CreateReception(ctx, reception)

	// Assert
	require.Error(err)
}
