package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPass(t *testing.T) {
	password := RandomPassword()

	hash, err := HashedPassword(password)

	require.NoError(t, err)

	require.NotEmpty(t, hash)

	err = CheckPassword(password, hash)

	require.NoError(t, err)

	wrongPassword := RandomPassword()

	err = CheckPassword(wrongPassword, hash)

	require.Error(t, err)
}
