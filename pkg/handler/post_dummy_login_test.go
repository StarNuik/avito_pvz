package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/starnuik/avito_pvz/pkg/pvztest"
	"github.com/stretchr/testify/require"
)

func Test_PostDummyLogin_Ok(t *testing.T) {
	// Arrange
	require := require.New(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	reqBody := "{\"role\":\"employee\"}"
	req, _ := http.NewRequest("POST", "/dummyLogin", strings.NewReader(reqBody))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(w.Code, 200)

	var result string
	err := json.Unmarshal(w.Body.Bytes(), &result)
	require.Nil(err)
	require.NotEmpty(result)
}

func Test_PostDummyLogin_BadJson_BadRequest(t *testing.T) {
	// Arrange
	require := require.New(t)

	app := pvztest.NewApp(t)

	w := httptest.NewRecorder()
	str := "{}"
	req, _ := http.NewRequest("POST", "/dummyLogin", strings.NewReader(str))

	// Act
	app.ServeHTTP(w, req)

	// Assert
	require.Equal(w.Code, 400)
}
