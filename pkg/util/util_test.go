package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test token read
func TestReadKubeTokenFromFile(t *testing.T) {
	require.FileExists(t, "./test_token")
	dat, err := ReadKubeTokenFromFile("./test_token")
	require.Equal(t, err, nil)
	digest := fmt.Sprintf("%x", md5.Sum([]byte(dat)))
	require.Equal(t, digest, "45168a1bf0ab611eab316af682c0ef91")
}

func TestLogin(t *testing.T) {
	token, err := ioutil.ReadFile("./test_token")
	require.Equal(t, err, nil)
	err = Login("http://localhost:8200/v1/auth/kubernetes/login", string(token))
	require.Equal(t, err, nil)
}

// Test write payload
func TestWriteSecret(t *testing.T) {
	WriteSecretPayload("./.env_test", `{"data":{"bar":"foo","foo":"bar"},"metadata":{"created_time":"2019-02-19T09:21:48.751378Z","deletion_time":"","destroyed":false,"version":1}}`)
	require.FileExists(t, "./.env_test")
	dat, err := ioutil.ReadFile("./.env_test")
	require.Equal(t, err, nil)
	digest := fmt.Sprintf("%x", md5.Sum(dat))
	require.Equal(t, digest, "b6d62d7836d92c66a390e26b60d99f1a")
}
