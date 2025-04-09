package usecase_test

import (
	"testing"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
	"github.com/stretchr/testify/require"
)

func Test_DummyLogin(t *testing.T) {
	// Arrange
	require := require.New(t)

	payload := token.Payload{
		UserRole: entity.UserRole(128),
	}

	usecase := usecase.New(nil, nil, nil)

	// Act
	result := usecase.DummyLogin(payload.UserRole)

	// Assert
	require.Equal(payload, result)
}
