package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArg(t *testing.T) {

	// Save original command line arguments
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Mock command line arguments
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	t.Run("ValidIndex", func(t *testing.T) {
		index := 1 // Use index 1 for the second element
		fallback := "defaultFallback"
		expectedValue := "arg2"

		result := GetArg(index, fallback)

		assert.Equal(t, expectedValue, result, "Expected value does not match")
	})

	t.Run("InvalidIndex", func(t *testing.T) {
		index := 3
		fallback := "defaultFallback"
		expectedFallback := fallback

		result := GetArg(index, fallback)

		assert.Equal(t, expectedFallback, result, "Fallback value does not match")
	})
}
