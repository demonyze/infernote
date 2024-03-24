package utils

import (
	"testing"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestFindNoteByName(t *testing.T) {
	type TestData struct {
		name     string
		expected model.Note
	}

	t.Run("Success", func(t *testing.T) {
		td := []TestData{
			{
				name:     "C",
				expected: data.Notes["C"],
			},
			{
				name:     "g#",
				expected: data.Notes["G#"],
			},
			{
				name:     "G#",
				expected: data.Notes["G#"],
			},
			{
				name:     "bb",
				expected: data.Notes["A#"],
			},
			{
				name:     "a ",
				expected: data.Notes["A"],
			},
			{
				name:     "		F#",
				expected: data.Notes["F#"],
			},
			{
				name:     " A#",
				expected: data.Notes["A#"],
			},
		}

		for _, testData := range td {
			note, err := FindNoteByName(testData.name)
			assert.NoError(t, err)
			assert.Equal(t, testData.expected, note)
		}
	})

	t.Run("Error", func(t *testing.T) {
		td := []TestData{
			{
				name:     "y",
				expected: model.Note{},
			},
			{
				name:     ".",
				expected: model.Note{},
			},
		}

		for _, testData := range td {
			note, err := FindNoteByName(testData.name)
			assert.Error(t, err)
			assert.Equal(t, testData.expected, note)
		}
	})

}
