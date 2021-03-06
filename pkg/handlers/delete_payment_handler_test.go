package handlers_test

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeletePaymentHandler_InvalidUUID(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("DELETE", "/1/payments/uuid", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid UUID"}`, rr.Body.String())
}

func TestDeletePaymentHandler_404(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("DELETE", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	require.Equal(t, `{"code":404,"message":"payment doesnt exist"}`, rr.Body.String())
}

func TestDeletePaymentHandler_OK(t *testing.T) {
	handler, repo, after := api.CreateTestingHandler(t)
	defer after()

	payment := payments.Payment{ID: "6a7d6b21-5cb7-4240-af3e-8dda39e65ff7"}
	err := repo.Save("6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", &payment)
	require.NoError(t, err)

	req, err := http.NewRequest("DELETE", "/1/payments/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNoContent, rr.Code)
}
