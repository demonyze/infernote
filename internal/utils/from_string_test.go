package utils

import (
	"testing"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNotesfromStringNames(t *testing.T) {
	type TestData struct {
		namesString   string
		expectedNotes []model.Note
		errorMessage  string
	}
	td := []TestData{
		{
			namesString:   "C#",
			expectedNotes: []model.Note{data.Notes["C#"]},
			errorMessage:  "Case 1: Single note failed",
		},
		{
			namesString: "F, G#, C, E, G",
			expectedNotes: []model.Note{
				data.Notes["F"],
				data.Notes["G#"],
				data.Notes["C"],
				data.Notes["E"],
				data.Notes["G"]},
			errorMessage: "Case 2: Multiple notes failed",
		},
		{
			namesString: "  A#, B, C#  ",
			expectedNotes: []model.Note{
				data.Notes["A#"],
				data.Notes["B"],
				data.Notes["C#"]},
			errorMessage: "Case 3: Notes with whitespaces failed",
		},
		{
			namesString:   "",
			expectedNotes: []model.Note{},
			errorMessage:  "Case 4: Empty string failed",
		},
		{
			namesString:   "...",
			expectedNotes: []model.Note{},
			errorMessage:  "Case 5: Wrong formatted string failed",
		},
	}

	for _, testData := range td {
		result := NotesfromStringNames(testData.namesString)
		assert.Equal(t, testData.expectedNotes, result, testData.errorMessage)
	}
}
