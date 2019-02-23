package payments

import (
	"github.com/google/uuid"
)

type PaymentIDGenerator interface {
	GenerateUniqueID() (uuid.UUID, error)
}

var _ PaymentIDGenerator = (*paymentIDGenerator)(nil)

type paymentIDGenerator struct{}

func NewPaymentIDGenerator() PaymentIDGenerator {
	return &paymentIDGenerator{}
}

func (*paymentIDGenerator) GenerateUniqueID() (uuid.UUID, error) {
	return uuid.NewRandom()
}
