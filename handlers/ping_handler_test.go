package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	handler := http.HandlerFunc(PingHandler)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	result := map[string]string{}

	err = json.Unmarshal(rr.Body.Bytes(), &result)
	require.NoError(t, err)
	require.Contains(t, result, "yes")
	require.Contains(t, result["yes"], "i_am")
}
