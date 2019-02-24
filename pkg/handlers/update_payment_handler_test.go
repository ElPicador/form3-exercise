package handlers_test

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdatePaymentHandler_InvalidUUID(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("PUT", "/1/payments/uuid", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid UUID"}`, rr.Body.String())
}

func TestUpdatePaymentHandler_EmptyBody(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("PUT", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"body of request must be a json object"}`, rr.Body.String())
}

func TestUpdatePaymentHandler_NotJSONBody(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("PUT", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", strings.NewReader("not a json"))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid json"}`, rr.Body.String())
}

func TestUpdatePaymentHandler_404(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("PUT", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", strings.NewReader("{}"))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	require.Equal(t, `{"code":404,"message":"payment doesnt exist"}`, rr.Body.String())
}

func TestUpdatePaymentHandler_ValidJSONBody(t *testing.T) {
	handler, repo, after := api.CreateTestingHandler(t)
	defer after()

	err := repo.Save("6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", &payments.Payment{
		ID:      "6a7d6b21-5cb7-4240-af3e-8dda39e65ff7",
		Version: 0,
	})
	require.NoError(t, err)

	req, err := http.NewRequest("PUT", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", strings.NewReader(`{"version":1}`))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNoContent, rr.Code)

	actual, err := repo.Get("6a7d6b21-5cb7-4240-af3e-8dda39e65ff7")
	expected := payments.Payment{
		ID:      "6a7d6b21-5cb7-4240-af3e-8dda39e65ff7",
		Version: 1,
	}
	require.NoError(t, err)
	require.EqualValues(t, expected, *actual)
}
