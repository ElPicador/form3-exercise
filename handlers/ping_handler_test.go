package handlers_test

import (
	"encoding/json"
	"github.com/ElPicador/form3-exercise/handlers"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestPingHandler(t *testing.T) {
	handler := http.HandlerFunc(handlers.PingHandler)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := handlers.ServeAndRecord(handler, req)

	require.Equal(t, http.StatusOK, rr.Code)
	result := map[string]string{}

	err = json.Unmarshal(rr.Body.Bytes(), &result)
	require.NoError(t, err)
	require.Contains(t, result, "yes")
	require.Contains(t, result["yes"], "i_am")
}
