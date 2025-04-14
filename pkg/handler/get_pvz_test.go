package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/middleware"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/stretchr/testify/require"
)

func Test_GetPvz_NoToken_Forbidden(t *testing.T) {
	// Arrange
	require := require.New(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pvz", nil)

	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(403, w.Code)
}

func Test_GetPvz_IncorrectRole_Forbidden(t *testing.T) {
	// Arrange
	require := require.New(t)

	tokenParser := token.NewParser()
	token, _ := tokenParser.Pack(token.Payload{
		UserRole: entity.UserRole(-1),
	})

	req, _ := http.NewRequest("GET", "/pvz", nil)
	req.Header.Add(middleware.AuthHeader, "Bearer "+token)

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(403, w.Code)
}

func Test_GetPvz_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	tokenParser := token.NewParser()
	token, _ := tokenParser.Pack(token.Payload{
		UserRole: entity.RoleEmployee,
	})

	startDate := time.Unix(0, 0)
	endDate := time.Unix(1, 0)
	startStr := url.QueryEscape(startDate.Format(time.RFC3339))
	endStr := url.QueryEscape(endDate.Format(time.RFC3339))
	query := fmt.Sprintf("/pvz?startDate=%s&endDate=%s", startStr, endStr)
	req, _ := http.NewRequest("GET", query, nil)
	req.Header.Add(middleware.AuthHeader, "Bearer "+token)

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(200, w.Code)
	require.Equal("[]", w.Body.String())
}
