package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/stretchr/testify/require"
)

func Test_LockPvz(t *testing.T) {
	// Arange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	pvz := entity.Pvz{Id: pvztest.NewUuid(t)}
	testRepo.CreatePvz(t, pvz)

	repo := pvztest.NewRepository(t)
	ctx := context.Background()

	table := []repository.DbLock{
		repository.LockAllowWrites, repository.LockNoWrites,
	}

	for _, row := range table {
		// Act
		tx, err := repo.LockPvz(ctx, pvz.Id, row)

		// Assert
		require.Nil(err, row)
		require.Nil(tx.Commit(ctx), row)
	}
}

func Test_LockPvz_NotFound(t *testing.T) {
	// Arange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	repo := pvztest.NewRepository(t)
	ctx := context.Background()

	table := []repository.DbLock{
		repository.LockAllowWrites, repository.LockNoWrites,
	}

	for _, row := range table {
		// Act
		tx, err := repo.LockPvz(ctx, uuid.Nil, row)

		// Assert
		require.Nil(tx, row)
		require.ErrorIs(err, entity.ErrNotFound, row)
	}
}
