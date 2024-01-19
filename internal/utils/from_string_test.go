package utils

import (
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNotesfromStringNames(t *testing.T) {
	// Test case 1: Single note
	namesString1 := "C"
	expectedNotes1 := []model.Note{{Name: "C"}}

	result1 := NotesfromStringNames(namesString1)
	assert.Equal(t, expectedNotes1, result1, "Case 1: Single note failed")

	// Test case 2: Multiple notes
	namesString2 := "C, D, E"
	expectedNotes2 := []model.Note{
		{Name: "C"},
		{Name: "D"},
		{Name: "E"},
	}

	result2 := NotesfromStringNames(namesString2)
	assert.Equal(t, expectedNotes2, result2, "Case 2: Multiple notes failed")

	// Test case 3: Notes with leading/trailing whitespaces
	namesString3 := "  A, B, C  "
	expectedNotes3 := []model.Note{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
	}

	result3 := NotesfromStringNames(namesString3)
	assert.Equal(t, expectedNotes3, result3, "Case 3: Notes with whitespaces failed")

	// Test case 4: Empty string
	namesString4 := ""
	expectedNotes4 := []model.Note{{Name: ""}}

	result4 := NotesfromStringNames(namesString4)
	assert.Equal(t, expectedNotes4, result4, "Case 4: Empty string failed")
}
