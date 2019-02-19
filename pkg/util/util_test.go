package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test write payload
func TestWriteSecret(t *testing.T) {
	WriteSecretPayload("./.env_test", `{"data":{"bar":"foo","foo":"bar"},"metadata":{"created_time":"2019-02-19T09:21:48.751378Z","deletion_time":"","destroyed":false,"version":1}}`)
	require.FileExists(t, "./.env_test")
	dat, err := ioutil.ReadFile("./.env_test")
	require.Equal(t, err, nil)
	digest := fmt.Sprintf("%x", md5.Sum(dat))
	require.Equal(t, digest, "81bafd0fd38daeafad2512e480a8f52d")
}
