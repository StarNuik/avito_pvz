package password_test

import (
	"testing"

	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func Test_HasherCompare(t *testing.T) {
	// Arrange
	require := require.New(t)

	table := []string{
		"",
		"a",
		"123",
		"abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz",
	}

	for idx, row := range table {
		bytes := []byte(row)
		len := min(72, len(bytes))
		bytes = bytes[:len]

		hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
		require.Nil(err)

		hasher := password.NewHasher()

		// Act
		result := hasher.Compare(row, hash)

		// Assert
		require.True(result, idx)
	}
}

func Test_HasherHash(t *testing.T) {
	// Arrange
	require := require.New(t)

	table := []string{
		"",
		"a",
		"123",
		"abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz",
	}

	for idx, row := range table {
		hasher := password.NewHasher()

		// Act
		result, err := hasher.Hash(row)

		// Assert
		require.Nil(err)

		err = bcrypt.CompareHashAndPassword(result, []byte(row))
		require.Nil(err, idx)
	}
}
