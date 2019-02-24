package payments_test

import (
	"github.com/ElPicador/form3-exercise/pkg/payments"
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

func TestGet(t *testing.T) {
	repo, after := payments.RepositoryForTests(t)
	defer after()

	id := "my-id"

	_, err := repo.Get(id)
	require.Error(t, err)

	payment := payments.Payment{ID: id}
	err = repo.Save(id, &payment)
	require.NoError(t, err)

	actual, err := repo.Get(id)
	require.NoError(t, err)
	require.EqualValues(t, payment, *actual)
}

func TestDelete(t *testing.T) {
	repo, after := payments.RepositoryForTests(t)
	defer after()

	id := "my-id"

	err := repo.Delete(id)
	require.Error(t, err)

	payment := payments.Payment{ID: id}
	err = repo.Save(id, &payment)
	require.NoError(t, err)

	err = repo.Delete(id)
	require.NoError(t, err)
}
