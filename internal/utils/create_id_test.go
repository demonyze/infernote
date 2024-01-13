package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateId(t *testing.T) {

	type TestData struct {
		key      string
		suffix   string
		expected string
	}

	testData := []TestData{
		{
			key:      "D",
			suffix:   "/F#",
			expected: "d/f#",
		},
		{
			key:      "A",
			suffix:   "Minor",
			expected: "aminor",
		},
	}

	for _, td := range testData {
		result := CreateId(td.key, td.suffix)
		assert.Equal(t, td.expected, result)
	}

}
