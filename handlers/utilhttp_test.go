package handlers

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestWrite500(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		Write500(w)
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := ServeAndRecord(handler, req)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	require.Equal(t, string(jsonBody500), rr.Body.String())
}

func TestWriteJSONMessage(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		WriteJSONMessage(w, http.StatusNotFound, "not found")
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := ServeAndRecord(handler, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	require.Equal(t, `{"code":404,"message":"not found"}`, rr.Body.String())
}

func TestWriteJSON(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		WriteJSON(w, http.StatusBadGateway, []byte(`{"hello":"world""}`))
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := ServeAndRecord(handler, req)

	require.Equal(t, http.StatusBadGateway, rr.Code)
	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	require.Equal(t, `{"hello":"world""}`, rr.Body.String())
}
