package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/app"
	"github.com/stretchr/testify/require"
)

func Test_GetPing_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	w := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/ping", nil)
	require.Nil(err)

	app, err := app.New()
	require.Nil(err)

	// Act
	app.ServeHTTP(w, request)

	// Assert
	require.Equal(w.Code, 200)
}
