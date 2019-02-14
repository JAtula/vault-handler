package secrets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRead a secret
func TestRead(t *testing.T) {
	callback, err := readSecret("/secret/data/test", "http://localhost:8200")
	require.Equal(t, err, nil)
	require.NotEmpty(t, callback)
}
