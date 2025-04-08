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

func Test_CreatePvz(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{
		Id:               uuid.Nil,
		RegistrationDate: time.Now().UTC(),
		City:             entity.PvzCity(0),
	}
	ctx := context.Background()
	repo := pvztest.NewRepository(t)

	// Act
	result, err := repo.CreatePvz(ctx, pvz)

	// Assert
	require.Nil(err)
	require.NotEqual(uuid.Nil, result.Id)

	require.Equal(pvz.City, result.City)

	// postgres timestamp accuracy is limited to micro-seconds
	microsec := pvz.RegistrationDate.UnixNano() / 1000
	date := time.UnixMicro(microsec).UTC()
	require.Equal(date, result.RegistrationDate)

}
