package util

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test write payload
func TestWriteSecret(t *testing.T) {
	writeSecretPayload("./test", "asdasd")
	require.FileExists(t, "./test")

	dat, err := ioutil.ReadFile("./test")
	require.Equal(t, err, nil)
	require.Equal(t, string(dat), "asdasd")
}
