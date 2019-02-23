package payments

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type testData struct {
	Data []Payment `json:"data"`
}

func TestPayment_JSONUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "data.json"))
	require.NoError(t, err)
	var testData testData

	err = json.Unmarshal(data, &testData)
	require.NoError(t, err)
}
