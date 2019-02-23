package payments

import (
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
