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
		Id: pvztest.NewUuid(t),
	}
	testRepo.CreatePvz(t, pvz)

	reception := entity.Reception{
		Id:       pvztest.NewUuid(t),
		PvzId:    pvz.Id,
		DateTime: time.Unix(1000, 0),
		Status:   entity.ReceptionStatus(128),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.CreateReception(ctx, reception)

	// Assert
	require.Nil(err)

	result := testRepo.GetReception(t, reception.Id)
	require.Equal(reception, result)
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
	err := repo.CreateReception(ctx, reception)

	// Assert
	require.ErrorIs(err, entity.ErrInternal)
}
