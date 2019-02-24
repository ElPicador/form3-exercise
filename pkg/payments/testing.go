package payments

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func RepositoryForTests(t *testing.T) (*Repository, func()) {
	dir, err := ioutil.TempDir("", "repository")
	require.NoError(t, err)

	return NewRepository(dir), func() {
		os.RemoveAll(dir) // clean up
	}
}

var _ PaymentIDGenerator = (*FixedPaymentIDGenerator)(nil)

type FixedPaymentIDGenerator struct{}

func (*FixedPaymentIDGenerator) GenerateUniqueID() (uuid.UUID, error) {
	return uuid.Parse("6a7d6b21-5cb7-4240-af3e-8dda39e65ff7")
}
