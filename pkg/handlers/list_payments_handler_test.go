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
	require.Equal(t, `{"payments":[{"type":"","id":"my-id-1","version":0,"organisation_id":"","attributes":{"amount":"","beneficiary_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"charges_information":{"bearer_code":"","sender_charges":null,"receiver_charges_amount":"","receiver_charges_currency":""},"currency":"","debtor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"end_to_end_reference":"","fx":{"contract_reference":"","exchange_rate":"","original_amount":"","original_currency":""},"numeric_reference":"","payment_id":"","payment_purpose":"","payment_scheme":"","payment_type":"","processing_date":"","reference":"","scheme_payment_sub_type":"","scheme_payment_type":"","sponsor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""}}},{"type":"","id":"my-id-2","version":0,"organisation_id":"","attributes":{"amount":"","beneficiary_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"charges_information":{"bearer_code":"","sender_charges":null,"receiver_charges_amount":"","receiver_charges_currency":""},"currency":"","debtor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""},"end_to_end_reference":"","fx":{"contract_reference":"","exchange_rate":"","original_amount":"","original_currency":""},"numeric_reference":"","payment_id":"","payment_purpose":"","payment_scheme":"","payment_type":"","processing_date":"","reference":"","scheme_payment_sub_type":"","scheme_payment_type":"","sponsor_party":{"account_name":"","account_number":"","account_number_code":"","account_type":0,"address":"","bank_id":"","bank_id_code":"","name":""}}}]}`, rr.Body.String())
}

