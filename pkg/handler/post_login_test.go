package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"github.com/starnuik/avito_pvz/pkg/dto"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_PostLogin_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	hasher := password.NewHasher()

	password := "password"
	passHash, _ := hasher.Hash(password)
	user := entity.User{
		Id:           uuid.Max,
		Email:        "email@email.com",
		PasswordHash: passHash,
		Role:         entity.RoleModerator,
	}

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)
	testRepo.CreateUser(t, user)

	reqDto := dto.PostLoginJSONBody{
		Email:    types.Email(user.Email),
		Password: password,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(200, w.Code)

	var resultToken string
	err := json.Unmarshal(w.Body.Bytes(), &resultToken)
	require.Nil(err)
	require.NotEmpty(resultToken)
}

func Test_PostLogin_IncorrectPassword_Unauthorized(t *testing.T) {
	// Arrange
	require := require.New(t)

	hasher := password.NewHasher()

	password := "password"
	passHash, _ := hasher.Hash(password)
	user := entity.User{
		Id:           uuid.Max,
		Email:        "email@email.com",
		PasswordHash: passHash,
		Role:         entity.RoleModerator,
	}

	testRepo := pvztest.NewTestRepository(t)
	testRepo.Clear(t)
	testRepo.CreateUser(t, user)

	reqDto := dto.PostLoginJSONBody{
		Email:    types.Email(user.Email),
		Password: "not" + password,
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(401, w.Code)
}

func Test_PostLogin_UserDoesntExist_Unauthorized(t *testing.T) {
	// Arrange
	require := require.New(t)

	pvztest.ClearRepository(t)

	reqDto := dto.PostLoginJSONBody{
		Email:    types.Email("doesnt@exist.com"),
		Password: "password",
	}
	reqBody, _ := json.Marshal(reqDto)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	app := pvztest.NewApp(t)

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(401, w.Code)
}
