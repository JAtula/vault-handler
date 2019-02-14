package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLogin the login
func TestLogin(t *testing.T) {
	callback, err := login("s.nEEaTPLjfEq4t9X4FWdIkAyQ", "http://localhost:8200")
	require.Equal(t, err, nil)
	require.NotEmpty(t, callback)
}
