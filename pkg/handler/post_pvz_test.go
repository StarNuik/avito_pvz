package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/middleware"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/stretchr/testify/require"
)

func Test_PostPvz_NoToken_Forbidden(t *testing.T) {
	// Arrange
	require := require.New(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/pvz", strings.NewReader("{}"))

	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(403, w.Code)
}

func Test_PostPvz_IncorrectRole_Forbidden(t *testing.T) {
	// Arrange
	require := require.New(t)

	tokenParser := token.NewParser()
	token, _ := tokenParser.Pack(token.Payload{
		UserRole: entity.RoleEmployee,
	})

	reqDto := dto.PVZ{
		Id:   &uuid.Max,
		City: dto.Казань,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/pvz", bytes.NewReader(reqBody))
	req.Header.Add(middleware.AuthHeader, "Bearer "+token)

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(403, w.Code)
}

func Test_PostPvz_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)

	tokenParser := token.NewParser()
	token, _ := tokenParser.Pack(token.Payload{
		UserRole: entity.RoleModerator,
	})

	reqDto := dto.PVZ{
		Id:   &uuid.Max,
		City: dto.Казань,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/pvz", bytes.NewReader(reqBody))
	req.Header.Add(middleware.AuthHeader, "Bearer "+token)

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(201, w.Code)

	require.Equal(1, testRepo.CountPvzs(t))
}
