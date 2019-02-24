package handlers_test

import (
	"github.com/ElPicador/form3-exercise/pkg/handlers"
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getHandler(t *testing.T) (http.Handler, *payments.Repository, func()) {
	repo, after := payments.RepositoryForTests(t)
	handler := http.Handler(handlers.NewGetPaymentHandler(repo))

	return handler, repo, after
}

func TestGetPaymentHandler_InvalidUUID(t *testing.T) {
	handler, _, after := getHandler(t)
	defer after()

	req, err := http.NewRequest("GET", "/uuid", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/{uuid}", handler)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, `{"code":400,"message":"invalid UUID"}`, rr.Body.String())
}

func TestGetPaymentHandler_404(t *testing.T) {
	handler, _, after := getHandler(t)
	defer after()

	req, err := http.NewRequest("GET", "/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/{uuid}", handler)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)
	require.Equal(t, `{"code":404,"message":"payment doesnt exist"}`, rr.Body.String())
}

func TestCreatePaymentHandler_OK(t *testing.T) {
	handler, repo, after := getHandler(t)
	defer after()

	payment := payments.Payment{ID: "6a7d6b21-5cb7-4240-af3e-8dda39e65ff7"}
	err := repo.Save("6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", &payment)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "/6a7d6b21-5cb7-4240-af3e-8dda39e65ff7", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/{uuid}", handler)
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, `{"payment":{"type":"","id":"6a7d6b21-5cb7-4240-af3e-8dda39e65ff7","version":0,"organisation_id":"","attributes":{"amount":"","beneficiary_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"charges_information":{"bearer_code":"","sender_charges":null,"receiver_charges_amount":"","receiver_charges_currency":""},"currency":"","debtor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"end_to_end_reference":"","fx":{"contract_reference":"","exchange_rate":"","original_amount":"","original_currency":""},"numeric_reference":"","payment_id":"","payment_purpose":"","payment_scheme":"","payment_type":"","processing_date":"","reference":"","scheme_payment_sub_type":"","scheme_payment_type":"","sponsor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""}}}}`, rr.Body.String())
}
