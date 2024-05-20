package precode

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenAllOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.NotEmpty(t, string(responseRecorder.Body.String()))
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=spb", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, 400, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, err := http.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	rRstring := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Len(t, rRstring, totalCount)
}
