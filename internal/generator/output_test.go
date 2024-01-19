package generator

import (
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestChordsAsMap(t *testing.T) {
	// Mock data for testing
	chordsDbImport := model.ChordsDbGuitarImport{
		Keys: []string{"C", "D"},
		Chords: map[string][]model.ChordsDbGuitarChord{
			"C": {
				{
					Key:       "C",
					Suffix:    "maj",
					Positions: []model.ChordsDbGuitarPosition{},
				},
			},
			"D": {
				{
					Key:       "D",
					Suffix:    "min",
					Positions: []model.ChordsDbGuitarPosition{},
				},
			},
		},
	}

	chordRocksImport := model.ChordRocksGuitarImport{
		"C": map[string]model.ChordRocksGuitarChord{
			"maj": {
				Notes: "C,E,G",
			},
		},
		"D": map[string]model.ChordRocksGuitarChord{
			"min": {
				Notes: "D,F,A",
			},
		},
	}

	params := CreateChordParams{
		chordsDbImport:   chordsDbImport,
		chordRocksImport: chordRocksImport,
	}

	// Test the ChordsAsMap function
	chordsMap := ChordsAsMap(params)

	assert.Len(t, chordsMap, 2, "Unexpected length of the chords map")
	assert.Contains(t, chordsMap, "dmin", "Key not found in the chords map")
}

func TestChordsAsArray(t *testing.T) {
	// Mock data for testing
	chordsDbImport := model.ChordsDbGuitarImport{
		Keys: []string{"C", "D"},
		Chords: map[string][]model.ChordsDbGuitarChord{
			"C": {
				{
					Key:       "C",
					Suffix:    "maj",
					Positions: []model.ChordsDbGuitarPosition{},
				},
			},
			"D": {
				{
					Key:       "D",
					Suffix:    "min",
					Positions: []model.ChordsDbGuitarPosition{},
				},
			},
		},
	}

	chordRocksImport := model.ChordRocksGuitarImport{
		"C": map[string]model.ChordRocksGuitarChord{
			"maj": {
				Notes: "C,E,G",
			},
		},
		"D": map[string]model.ChordRocksGuitarChord{
			"min": {
				Notes: "D,F,A",
			},
		},
	}

	params := CreateChordParams{
		chordsDbImport:   chordsDbImport,
		chordRocksImport: chordRocksImport,
	}

	// Test the ChordsAsArray function
	chordsArray := ChordsAsArray(params)

	assert.Len(t, chordsArray, 2, "Unexpected length of the chords array")
	assert.Equal(t, "dmin", chordsArray[1].Id, "Chord does not match expectations")
}
