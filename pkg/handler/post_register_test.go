package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_PostRegister_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	pvztest.ClearRepository(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	reqDto := dto.PostRegisterJSONBody{
		Email:    "abc@email.com",
		Password: "weak password",
		Role:     dto.Moderator,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(reqBody)))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(201, w.Code)
}

func Test_PostRegister_DuplicateEmail_BadRequest(t *testing.T) {
	// Arrange
	require := require.New(t)

	pvztest.ClearRepository(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	reqDto := dto.PostRegisterJSONBody{
		Email:    "abc@email.com",
		Password: "weak password",
		Role:     dto.Moderator,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(reqBody)))

	app.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/register", strings.NewReader(string(reqBody)))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(400, w.Code)
}

func Test_PostRegister_EmptyEmail_BadRequest(t *testing.T) {
	// Arrange
	require := require.New(t)

	pvztest.ClearRepository(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	reqDto := dto.PostRegisterJSONBody{
		Email:    "",
		Password: "weak password",
		Role:     dto.Moderator,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(reqBody)))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(400, w.Code)
}

func Test_PostRegister_EmptyPassword_BadRequest(t *testing.T) {
	// Arrange
	require := require.New(t)

	pvztest.ClearRepository(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	reqDto := dto.PostRegisterJSONBody{
		Email:    "abc@email.com",
		Password: "",
		Role:     dto.Moderator,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(reqBody)))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(400, w.Code)
}
