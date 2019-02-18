package secrets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRead a secret
func TestRead(t *testing.T) {
	callback, err := ReadSecret("/secret/data/test", os.Getenv("VAULT_ADDR"))
	require.Equal(t, err, nil)
	require.NotEqual(t, callback, "404")
}
