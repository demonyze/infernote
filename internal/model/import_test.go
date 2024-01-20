package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChordsDbGuitarImport(t *testing.T) {
	dataImport := ChordsDbGuitarImport{
		Keys: []string{"C", "D"},
		Chords: map[string][]ChordsDbGuitarChord{
			"C": {
				{
					Key:       "C",
					Suffix:    "maj",
					Positions: []ChordsDbGuitarPosition{},
				},
			},
			"D": {
				{
					Key:       "D",
					Suffix:    "min",
					Positions: []ChordsDbGuitarPosition{},
				},
			},
		},
	}

	assert.Len(t, dataImport.Keys, 2, "Unexpected length of Keys")
	assert.Contains(t, dataImport.Keys, "C", "Key 'C' not found in Keys")
	assert.Contains(t, dataImport.Keys, "D", "Key 'D' not found in Keys")

	assert.Len(t, dataImport.Chords, 2, "Unexpected length of Chords")
	assert.Contains(t, dataImport.Chords, "C", "Key 'C' not found in Chords")
	assert.Contains(t, dataImport.Chords, "D", "Key 'D' not found in Chords")
}

func TestChordsDbGuitarChord(t *testing.T) {
	chord := ChordsDbGuitarChord{
		Key:       "C",
		Suffix:    "maj",
		Positions: []ChordsDbGuitarPosition{},
	}

	assert.Equal(t, "C", chord.Key, "Unexpected Key value")
	assert.Equal(t, "maj", chord.Suffix, "Unexpected Suffix value")
	assert.Len(t, chord.Positions, 0, "Unexpected length of Positions")
}

func TestChordsDbGuitarPosition(t *testing.T) {
	position := ChordsDbGuitarPosition{
		Frets:    []int{1, 2, 3},
		Fingers:  []int{1, 2, 3},
		Barres:   []int{1},
		Capo:     true,
		BaseFret: 1,
		Midi:     []int{60, 61, 62},
	}

	assert.Equal(t, []int{1, 2, 3}, position.Frets, "Unexpected Frets value")
	assert.Equal(t, []int{1, 2, 3}, position.Fingers, "Unexpected Fingers value")
	assert.Equal(t, []int{1}, position.Barres, "Unexpected Barres value")
	assert.True(t, position.Capo, "Unexpected Capo value")
	assert.Equal(t, 1, position.BaseFret, "Unexpected BaseFret value")
	assert.Equal(t, []int{60, 61, 62}, position.Midi, "Unexpected Midi value")
}
