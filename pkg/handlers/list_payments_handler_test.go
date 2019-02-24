package handlers_test

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPaymentsHandler_NoPayments(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("GET", "/1/payments", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, `{"payments":[]}`, rr.Body.String())
}

func TestListPaymentsHandler_2Payments(t *testing.T) {
	handler, repo, after := api.CreateTestingHandler(t)
	defer after()

	err := repo.Save("my-id-1", &payments.Payment{ID: "my-id-1"})
	require.NoError(t, err)
	err = repo.Save("my-id-2", &payments.Payment{ID: "my-id-2"})
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "/1/payments", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, `{"payments":[{"id":"my-id-1"},{"id":"my-id-2"}]}`, rr.Body.String())
}

