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
		languageImports:  []model.LanguageImport{},
	}

	// Test the ChordsAsMap function
	chordsMap := ChordsAsMap(params)

	assert.Len(t, chordsMap, 2, "Unexpected length of the chords map")
	assert.Contains(t, chordsMap, "dmin", "Key not found in the chords map")
}

func TestMapChordsDbNote(t *testing.T) {
	tests := []struct {
		note     model.Note
		expected string
	}{
		{model.Note{Name: "A#", AlternativeName: "Bb"}, "Bb"},
		{model.Note{Name: "C#", AlternativeName: "Db"}, "Csharp"},
		{model.Note{Name: "D#", AlternativeName: "Eb"}, "Eb"},
		{model.Note{Name: "F#", AlternativeName: "Gb"}, "Fsharp"},
		{model.Note{Name: "G#", AlternativeName: "Ab"}, "Ab"},
		{model.Note{Name: "C"}, "C"},
		{model.Note{Name: "F"}, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.note.Name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mapChordsDbNote(tt.note))
		})
	}
}

func TestTypes(t *testing.T) {
	t.Run("Single language", func(t *testing.T) {
		imports := []model.LanguageImport{
			{
				Language: "en",
				Types: map[string]string{
					"major": "Major",
					"minor": "Minor",
					"5":     "5th",
					"7":     "7th",
				},
			},
		}

		expected := map[string]model.ChordType{
			"major": {Id: "major", Lang: map[string]string{"en": "Major"}},
			"minor": {Id: "minor", Lang: map[string]string{"en": "Minor"}},
			"5":     {Id: "5", Lang: map[string]string{"en": "5th"}},
			"7":     {Id: "7", Lang: map[string]string{"en": "7th"}},
		}

		result := Types(imports)
		assert.Equal(t, expected, result)
	})

	t.Run("Multiple languages", func(t *testing.T) {
		imports := []model.LanguageImport{
			{
				Language: "en",
				Types:    map[string]string{"major": "Major", "minor": "Minor"},
			},
			{
				Language: "de",
				Types:    map[string]string{"major": "Dur", "minor": "Moll"},
			},
		}

		result := Types(imports)

		assert.Equal(t, "major", result["major"].Id)
		assert.Equal(t, "Major", result["major"].Lang["en"])
		assert.Equal(t, "Dur", result["major"].Lang["de"])
		assert.Equal(t, "Minor", result["minor"].Lang["en"])
		assert.Equal(t, "Moll", result["minor"].Lang["de"])
	})
}
