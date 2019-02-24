package api

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
	"net/http"
	"testing"
)

func CreateTestingHandler(t *testing.T) (http.Handler, *payments.Repository, func()) {
	repo, after := payments.RepositoryForTests(t)
	handler := New(repo, &payments.FixedPaymentIDGenerator{})

	return handler.Handler(), repo, after
}
