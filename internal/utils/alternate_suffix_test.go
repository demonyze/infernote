package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternateSuffx(t *testing.T) {
	type TestData struct {
		input  string
		result string
	}

	testData := []TestData{
		{input: "major", result: ""},
		{input: "minor", result: "m"},
		{input: "maj7", result: "maj7"},
	}

	for _, td := range testData {
		suffix := AlternateSuffix(td.input)
		assert.Equal(t, td.result, suffix)
	}
}
