package payments_test

import (
	"github.com/ElPicador/form3-exercise/payments"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	dir, err := ioutil.TempDir("", "repository")
	require.NoError(t, err)

	defer os.RemoveAll(dir) // clean up

	repo := payments.NewRepository(dir)

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
