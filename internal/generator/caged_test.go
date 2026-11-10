package generator

import (
	"fmt"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

type chordFormula struct {
	allowedIntervals  []int
	requiredIntervals []int
}

var chordFormulasBySuffix = map[string]chordFormula{
	"major": {[]int{0, 4, 7}, []int{0, 4}},
	"minor": {[]int{0, 3, 7}, []int{0, 3}},
	"7":     {[]int{0, 4, 7, 10}, []int{0, 4, 10}},
	"m7":    {[]int{0, 3, 7, 10}, []int{0, 3, 10}},
	"maj7":  {[]int{0, 4, 7, 11}, []int{0, 4, 11}},
}

var allRoots = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

func intervalsAboveRoot(midi []int, rootPitchClass int) map[int]bool {
	intervals := map[int]bool{}
	for _, note := range midi {
		intervals[semitonesUpFrom(rootPitchClass, note%semitones)] = true
	}
	return intervals
}

func TestEveryCagedPositionIsAPlayableVoicingOfItsChord(t *testing.T) {
	assert.Len(t, cagedShapesBySuffix, len(chordFormulasBySuffix))
	for suffix, formula := range chordFormulasBySuffix {
		for _, root := range allRoots {
			t.Run(fmt.Sprintf("%s %s", root, suffix), func(t *testing.T) {
				positions := CagedPositions(root, suffix)
				assert.Len(t, positions, 5)

				rootPitchClass := pitchClasses[root]
				for i, position := range positions {
					shape := cagedShapesBySuffix[suffix][i].name
					assert.Len(t, position.Frets, stringCount, shape)
					assert.Len(t, position.Fingers, stringCount, shape)
					assert.GreaterOrEqual(t, position.BaseFret, 1, shape)

					assert.Equal(t, rootPitchClass, position.Midi[0]%semitones, "%s shape: bass must be the root", shape)
					intervals := intervalsAboveRoot(position.Midi, rootPitchClass)
					for interval := range intervals {
						assert.Contains(t, formula.allowedIntervals, interval, "%s shape: unexpected interval", shape)
					}
					for _, interval := range formula.requiredIntervals {
						assert.True(t, intervals[interval], "%s shape: missing interval %d", shape, interval)
					}

					for s, fret := range position.Frets {
						assert.LessOrEqual(t, fret, diagramHeight, "%s shape: fret outside diagram", shape)
						if fret > openString {
							assert.NotZero(t, position.Fingers[s], "%s shape: fretted string without finger", shape)
						} else {
							assert.Zero(t, position.Fingers[s], "%s shape: finger on open or muted string", shape)
						}
					}
				}
			})
		}
	}
}

func TestCagedPositionsReturnsNilForUnknownRootOrSuffix(t *testing.T) {
	assert.Nil(t, CagedPositions("H", "major"))
	assert.Nil(t, CagedPositions("C", "9"))
}

func TestCagedPositionsForCMajor(t *testing.T) {
	positionsByShape := map[string]model.ChordsDbGuitarPosition{}
	for i, position := range CagedPositions("C", "major") {
		positionsByShape[cagedShapesBySuffix["major"][i].name] = position
	}

	t.Run("C shape is the open chord", func(t *testing.T) {
		assert.Equal(t, model.ChordsDbGuitarPosition{
			Frets: []int{-1, 3, 2, 0, 1, 0}, Fingers: []int{0, 3, 2, 0, 1, 0},
			Barres: []int{}, Capo: false, BaseFret: 1, Midi: []int{48, 52, 55, 60, 64},
		}, positionsByShape["C"])
	})
	t.Run("A shape fits the diagram from the nut", func(t *testing.T) {
		assert.Equal(t, model.ChordsDbGuitarPosition{
			Frets: []int{-1, 3, 5, 5, 5, 3}, Fingers: []int{0, 1, 2, 3, 4, 1},
			Barres: []int{3}, Capo: true, BaseFret: 1, Midi: []int{48, 55, 60, 64, 67},
		}, positionsByShape["A"])
	})
	t.Run("G shape has a partial barre", func(t *testing.T) {
		assert.Equal(t, model.ChordsDbGuitarPosition{
			Frets: []int{4, 3, 1, 1, 1, 4}, Fingers: []int{3, 2, 1, 1, 1, 4},
			Barres: []int{1}, Capo: false, BaseFret: 5, Midi: []int{48, 52, 55, 60, 64, 72},
		}, positionsByShape["G"])
	})
	t.Run("E shape has a full barre", func(t *testing.T) {
		assert.Equal(t, model.ChordsDbGuitarPosition{
			Frets: []int{1, 3, 3, 2, 1, 1}, Fingers: []int{1, 3, 4, 2, 1, 1},
			Barres: []int{1}, Capo: true, BaseFret: 8, Midi: []int{48, 55, 60, 64, 67, 72},
		}, positionsByShape["E"])
	})
	t.Run("D shape has no barre", func(t *testing.T) {
		assert.Equal(t, model.ChordsDbGuitarPosition{
			Frets: []int{-1, -1, 1, 3, 4, 3}, Fingers: []int{0, 0, 1, 2, 4, 3},
			Barres: []int{}, Capo: false, BaseFret: 10, Midi: []int{60, 67, 72, 76},
		}, positionsByShape["D"])
	})
}

func TestSemitonesUpFrom(t *testing.T) {
	assert.Equal(t, 0, semitonesUpFrom(4, 4))
	assert.Equal(t, 3, semitonesUpFrom(9, 0))
	assert.Equal(t, 10, semitonesUpFrom(2, 0))
}

func TestShiftFretsKeepsMutedStrings(t *testing.T) {
	assert.Equal(t, perString{-1, 5, 7, 7, 7, 5}, shiftFrets(perString{-1, 0, 2, 2, 2, 0}, 5))
}

func TestDiagramBaseFret(t *testing.T) {
	assert.Equal(t, 1, diagramBaseFret(perString{-1, 3, 5, 5, 5, 3}))
	assert.Equal(t, 1, diagramBaseFret(perString{0, 0, 0, 0, 0, 0}))
	assert.Equal(t, 8, diagramBaseFret(perString{8, 10, 10, 9, 8, 8}))
}

func TestIndexFingerBarre(t *testing.T) {
	t.Run("without index finger", func(t *testing.T) {
		barreFret, barredStrings := indexFingerBarre(perString{-1, 3, 2, 0, 1, 0}, perString{0, 3, 2, 0, 0, 0})
		assert.Equal(t, 0, barreFret)
		assert.Empty(t, barredStrings)
	})
	t.Run("index finger on different frets", func(t *testing.T) {
		barreFret, barredStrings := indexFingerBarre(perString{3, 4, -1, -1, -1, -1}, perString{1, 1, 0, 0, 0, 0})
		assert.Equal(t, 3, barreFret)
		assert.Equal(t, []int{0}, barredStrings)
	})
	t.Run("index finger across several strings", func(t *testing.T) {
		barreFret, barredStrings := indexFingerBarre(perString{-1, 5, 7, 7, 7, 5}, perString{0, 1, 2, 3, 4, 1})
		assert.Equal(t, 5, barreFret)
		assert.Equal(t, []int{1, 5}, barredStrings)
	})
}

func TestLowestSoundingString(t *testing.T) {
	assert.Equal(t, 0, lowestSoundingString(perString{0, 2, 2, 1, 0, 0}))
	assert.Equal(t, 2, lowestSoundingString(perString{-1, -1, 0, 2, 3, 2}))
	assert.Equal(t, mutedString, lowestSoundingString(perString{-1, -1, -1, -1, -1, -1}))
}

func TestFretsOfConvertsDiagramFretsToNeckFrets(t *testing.T) {
	assert.Equal(t, perString{-1, 3, 5, 5, 5, 3}, fretsOf(model.ChordsDbGuitarPosition{
		Frets: []int{-1, 1, 3, 3, 3, 1}, BaseFret: 3,
	}))
	assert.Equal(t, perString{-1, 0, 2, 2, 2, 5}, fretsOf(model.ChordsDbGuitarPosition{
		Frets: []int{-1, 0, 1, 1, 1, 4}, BaseFret: 2,
	}))
	assert.Equal(t, perString{0, 2, -1, -1, -1, -1}, fretsOf(model.ChordsDbGuitarPosition{
		Frets: []int{0, 2}, BaseFret: 1,
	}))
}

func TestAppendNewVoicingsKeepsExistingOrderAndSortsAdded(t *testing.T) {
	existingOpen := model.ChordsDbGuitarPosition{Frets: []int{-1, 3, 2, 0, 1, 0}, BaseFret: 1}
	existingHigh := model.ChordsDbGuitarPosition{Frets: []int{1, 3, 3, 2, 1, 1}, BaseFret: 8}
	existingA := model.ChordsDbGuitarPosition{Frets: []int{-1, 1, 3, 3, 3, 1}, BaseFret: 3}

	sameAsExistingA := model.ChordsDbGuitarPosition{Frets: []int{-1, 3, 5, 5, 5, 3}, BaseFret: 1}
	newHigh := model.ChordsDbGuitarPosition{Frets: []int{-1, -1, 1, 3, 4, 3}, BaseFret: 10}
	newMid := model.ChordsDbGuitarPosition{Frets: []int{4, 3, 1, 1, 1, 4}, BaseFret: 5}

	merged := AppendNewVoicings(
		[]model.ChordsDbGuitarPosition{existingOpen, existingHigh, existingA},
		[]model.ChordsDbGuitarPosition{sameAsExistingA, newHigh, newMid, newHigh},
	)

	assert.Equal(t, []model.ChordsDbGuitarPosition{
		existingOpen, existingHigh, existingA, newMid, newHigh,
	}, merged)
}

func TestAppendNewVoicingsSkipsVoicingsThatDifferOnlyByMutedStrings(t *testing.T) {
	partialG := model.ChordsDbGuitarPosition{Frets: []int{-1, -1, 1, 1, 1, 4}, BaseFret: 5}
	fullA := model.ChordsDbGuitarPosition{Frets: []int{-1, 1, 1, 3, 4, 3}, BaseFret: 4}

	fullG := model.ChordsDbGuitarPosition{Frets: []int{4, 3, 1, 1, 1, 4}, BaseFret: 5}
	partialA := model.ChordsDbGuitarPosition{Frets: []int{-1, -1, 1, 3, 4, 3}, BaseFret: 4}

	merged := AppendNewVoicings(
		[]model.ChordsDbGuitarPosition{fullA, partialG},
		[]model.ChordsDbGuitarPosition{fullG, partialA},
	)

	assert.Equal(t, []model.ChordsDbGuitarPosition{fullA, partialG}, merged)
}

func TestAppendNewVoicingsWithNothingGenerated(t *testing.T) {
	existing := []model.ChordsDbGuitarPosition{{Frets: []int{0, 2, 2, 0, 0, 0}, BaseFret: 1}}
	assert.Equal(t, existing, AppendNewVoicings(existing, nil))
	assert.Empty(t, AppendNewVoicings(nil, nil))
}

func TestDifferOnlyByMutedStrings(t *testing.T) {
	tests := []struct {
		name     string
		a, b     perString
		expected bool
	}{
		{"identical", perString{0, 2, 2, 1, 0, 0}, perString{0, 2, 2, 1, 0, 0}, true},
		{"a mutes strings b plays", perString{-1, -1, 5, 5, 5, 8}, perString{8, 7, 5, 5, 5, 8}, true},
		{"b mutes strings a plays", perString{-1, 4, 4, 6, 7, 6}, perString{-1, -1, 4, 6, 7, 6}, true},
		{"different fret on a shared string", perString{-1, 3, 2, 0, 1, 0}, perString{-1, 3, 2, 0, 1, 3}, false},
		{"each plays a string the other mutes", perString{0, 2, 2, -1, -1, -1}, perString{-1, 2, 2, 1, -1, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, differOnlyByMutedStrings(tt.a, tt.b))
		})
	}
}
