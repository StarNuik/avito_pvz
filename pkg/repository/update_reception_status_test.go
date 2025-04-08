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

func Test_UpdateReceptionStatus_ReceptionExists(t *testing.T) {
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
		Id:       pvztest.NewUuid(t),
		PvzId:    pvz.Id,
		DateTime: time.Now().UTC(),
		Status:   entity.ReceptionStatus(0),
	}
	testRepo.CreateReception(t, reception)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.UpdateReceptionStatus(ctx, reception.Id, entity.ReceptionStatus(1))

	// Assert
	require.Nil(err)
	require.Equal(entity.ReceptionStatus(1), result.Status)

	require.Equal(reception.Id, result.Id)
	pvztest.RequireEqualTime(t, reception.DateTime, result.DateTime)
}

func Test_UpdateReceptionStatus_NoReception(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	_, err := repo.UpdateReceptionStatus(ctx, uuid.Nil, entity.ReceptionStatus(1))

	// Assert
	require.Error(err)
}
