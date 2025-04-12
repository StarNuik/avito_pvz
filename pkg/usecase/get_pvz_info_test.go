package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/mocks"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_GetPvzInfo_AllOptions(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	pvzInfo := entity.PvzInfo{}
	startDate := time.Unix(500, 0)
	endDate := time.Unix(1000, 0)
	page := 5
	limit := 5
	offset := (page - 1) * limit

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		GetPvzInfo(gomock.Any(), startDate, endDate, limit, offset).
		Return(pvzInfo, nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()

	// Act
	result, err := usecase.GetPvzInfo(ctx, startDate, endDate, &page, &limit)

	// Assert
	require.Nil(err)
	require.Equal(pvzInfo, result)
}

func Test_GetPvzInfo_NoOptions(t *testing.T) {
	// Arrange
	require := require.New(t)
	ctrl := gomock.NewController(t)

	pvzInfo := entity.PvzInfo{}
	startDate := time.Unix(500, 0)
	endDate := time.Unix(1000, 0)
	limit := 10
	offset := 0

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		GetPvzInfo(gomock.Any(), startDate, endDate, limit, offset).
		Return(pvzInfo, nil)

	usecase := usecase.New(repo, nil, nil)

	ctx := context.Background()

	// Act
	result, err := usecase.GetPvzInfo(ctx, startDate, endDate, nil, nil)

	// Assert
	require.Nil(err)
	require.Equal(pvzInfo, result)
}
