package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_CreatePvz(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{
		Id:               pvztest.NewUuid(t),
		RegistrationDate: time.Unix(1000, 0),
		City:             entity.PvzCity(128),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	err := repo.CreatePvz(ctx, pvz)

	// Assert
	require.Nil(err)

	result := testRepo.GetPvz(t, pvz.Id)
	require.Equal(pvz, result)
}
