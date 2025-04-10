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

func Test_GetLastReception(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{}
	testRepo.CreatePvz(t, pvz)

	receptions := []entity.Reception{
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(100, 0), PvzId: pvz.Id},
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(200, 0), PvzId: pvz.Id},
		{Id: pvztest.NewUuid(t), DateTime: time.Unix(300, 0), PvzId: pvz.Id},
	}
	for _, r := range receptions {
		testRepo.CreateReception(t, r)
	}

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.GetLastReception(ctx, pvz.Id)

	// Assert
	require.Nil(err)
	require.Equal(receptions[2], result)
}

func Test_GetLastReception_NotFound(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	_, err := repo.GetLastReception(ctx, uuid.Nil)

	// Assert
	require.ErrorIs(err, entity.ErrNotFound)
}
