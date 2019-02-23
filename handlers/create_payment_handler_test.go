package handlers_test

import (
	"github.com/ElPicador/form3-exercise/handlers"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestCreatePaymentHandler_EmptyBody(t *testing.T) {
	handler := http.HandlerFunc(handlers.CreatePaymentHandler)

	req, err := http.NewRequest("POST", "/", nil)
	require.NoError(t, err)

	rr := handlers.ServeAndRecord(handler, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"body of request must be a json object"}`, rr.Body.String())
}

func TestCreatePaymentHandler_NotJSONBody(t *testing.T) {
	handler := http.HandlerFunc(handlers.CreatePaymentHandler)

	req, err := http.NewRequest("POST", "/", strings.NewReader("not a json"))
	require.NoError(t, err)

	rr := handlers.ServeAndRecord(handler, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid json"}`, rr.Body.String())
}

func TestCreatePaymentHandler_ValidJSONBody(t *testing.T) {
	handler := http.HandlerFunc(handlers.CreatePaymentHandler)

	req, err := http.NewRequest("POST", "/", strings.NewReader("{}"))
	require.NoError(t, err)

	rr := handlers.ServeAndRecord(handler, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	require.Equal(t, `{"payment_id":"uuidv4"}`, rr.Body.String())
}
