package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test token read
func TestReadTokenFromFile(t *testing.T) {
	require.FileExists(t, "./test_token")
	dat, err := ReadTokenFromFile("./test_token")
	require.Equal(t, err, nil)
	digest := fmt.Sprintf("%x", md5.Sum([]byte(dat)))
	require.Equal(t, digest, "383aad497c96fdd9453ada3877fd748a")
}

//Test setting token as environment variable
func TestSetVaultToken(t *testing.T) {
	err := SetVaultToken("asd")
	require.Equal(t, err, nil)
	require.NotEmpty(t, os.Getenv("VAULT_TOKEN"))
	require.Equal(t, os.Getenv("VAULT_TOKEN"), "asd")
}

// Test write payload
func TestWriteSecret(t *testing.T) {
	WriteSecretPayload("./.env_test", `{"data":{"bar":"foo","foo":"bar"},"metadata":{"created_time":"2019-02-19T09:21:48.751378Z","deletion_time":"","destroyed":false,"version":1}}`)
	require.FileExists(t, "./.env_test")
	dat, err := ioutil.ReadFile("./.env_test")
	require.Equal(t, err, nil)
	digest := fmt.Sprintf("%x", md5.Sum(dat))
	require.Equal(t, digest, "81bafd0fd38daeafad2512e480a8f52d")
}
