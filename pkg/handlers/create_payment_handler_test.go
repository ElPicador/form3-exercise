package handlers_test

import (
	"github.com/ElPicador/form3-exercise/pkg/api"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestCreatePaymentHandler_EmptyBody(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("POST", "/1/payments", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"body of request must be a json object"}`, rr.Body.String())
}

func TestCreatePaymentHandler_NotJSONBody(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("POST", "/1/payments", strings.NewReader("not a json"))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid json"}`, rr.Body.String())
}

func TestCreatePaymentHandler_ValidJSONBody(t *testing.T) {
	handler, _, after := api.CreateTestingHandler(t)
	defer after()

	req, err := http.NewRequest("POST", "/1/payments", strings.NewReader("{}"))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)

	require.Equal(t, `{"payment_id":"6a7d6b21-5cb7-4240-af3e-8dda39e65ff7"}`, rr.Body.String())
}
