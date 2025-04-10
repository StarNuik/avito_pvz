package repository_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_GetPvzInfo(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	info := pvztest.FilledPvzInfo(t)
	insertPvzInfo(t, testRepo, info)

	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.GetPvzInfo(ctx, time.Unix(0, 0), time.Unix(100000, 0), math.MaxInt, 0)

	// Assert
	require.Nil(err)
	require.Equal(info, result)
}

func insertPvzInfo(t *testing.T, repo *pvztest.TestRepository, info entity.PvzInfo) {
	for _, pvz := range info.Pvzs {
		repo.CreatePvz(t, pvz)
	}
	for _, r := range info.Receptions {
		repo.CreateReception(t, r)
	}
	for _, p := range info.Products {
		repo.CreateProduct(t, p)
	}
}
