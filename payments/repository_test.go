package payments_test

import (
	"github.com/ElPicador/form3-exercise/payments"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExists(t *testing.T) {
	repo, after := payments.RepositoryForTests(t)
	defer after()

	id := "my-id"

	exists, err := repo.Exists(id)
	require.NoError(t, err)
	require.False(t, exists)

	err = repo.Save(id, &payments.Payment{})
	require.NoError(t, err)

	exists, err = repo.Exists(id)
	require.NoError(t, err)
	require.True(t, exists)
}
