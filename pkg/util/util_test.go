package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLogin the login
func TestReadToken(t *testing.T) {
	cdFalse, err := readToken("")
	require.Equal(t, err, nil)
	require.NotEqual(t, cdFalse, "Couldn't find env.")

	cdTrue, err := readToken("asdasd")
	require.Equal(t, err, nil)
	require.Equal(t, cdTrue, "Don't read default token")
}
