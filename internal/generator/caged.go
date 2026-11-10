package generator

import (
	"sort"

	"github.com/demonyze/infernote/internal/model"
)

const (
	stringCount   = 6
	mutedString   = -1
	openString    = 0
	indexFinger   = 1
	semitones     = 12
	diagramHeight = 5
)

type perString [stringCount]int

var standardTuningMidi = perString{40, 45, 50, 55, 59, 64}

var pitchClasses = map[string]int{
	"C": 0, "C#": 1, "Db": 1, "D": 2, "D#": 3, "Eb": 3, "E": 4, "F": 5,
	"F#": 6, "Gb": 6, "G": 7, "G#": 8, "Ab": 8, "A": 9, "A#": 10, "Bb": 10, "B": 11,
}

type voicing struct {
	frets   perString
	fingers perString
}

// A cagedShape is named after the open chord it comes from. Its movable frets
// are counted from the barre that replaces the nut, so they can be shifted to any root.
type cagedShape struct {
	name    string
	movable voicing
	open    voicing
}

var cagedShapesBySuffix = map[string][]cagedShape{
	"major": {
		{"C", voicing{perString{-1, 3, 2, 0, 1, 0}, perString{0, 4, 3, 1, 2, 1}}, voicing{perString{-1, 3, 2, 0, 1, 0}, perString{0, 3, 2, 0, 1, 0}}},
		{"A", voicing{perString{-1, 0, 2, 2, 2, 0}, perString{0, 1, 2, 3, 4, 1}}, voicing{perString{-1, 0, 2, 2, 2, 0}, perString{0, 0, 1, 2, 3, 0}}},
		{"G", voicing{perString{3, 2, 0, 0, 0, 3}, perString{3, 2, 1, 1, 1, 4}}, voicing{perString{3, 2, 0, 0, 0, 3}, perString{2, 1, 0, 0, 0, 3}}},
		{"E", voicing{perString{0, 2, 2, 1, 0, 0}, perString{1, 3, 4, 2, 1, 1}}, voicing{perString{0, 2, 2, 1, 0, 0}, perString{0, 2, 3, 1, 0, 0}}},
		{"D", voicing{perString{-1, -1, 0, 2, 3, 2}, perString{0, 0, 1, 2, 4, 3}}, voicing{perString{-1, -1, 0, 2, 3, 2}, perString{0, 0, 0, 1, 3, 2}}},
	},
	"minor": {
		{"C", voicing{perString{-1, 3, 1, 0, 1, -1}, perString{0, 4, 2, 1, 3, 0}}, voicing{perString{-1, 3, 1, 0, 1, -1}, perString{0, 3, 1, 0, 2, 0}}},
		{"A", voicing{perString{-1, 0, 2, 2, 1, 0}, perString{0, 1, 3, 4, 2, 1}}, voicing{perString{-1, 0, 2, 2, 1, 0}, perString{0, 0, 2, 3, 1, 0}}},
		{"G", voicing{perString{3, 1, 0, 0, 3, -1}, perString{3, 2, 1, 1, 4, 0}}, voicing{perString{3, 1, 0, 0, 3, -1}, perString{2, 1, 0, 0, 3, 0}}},
		{"E", voicing{perString{0, 2, 2, 0, 0, 0}, perString{1, 3, 4, 1, 1, 1}}, voicing{perString{0, 2, 2, 0, 0, 0}, perString{0, 2, 3, 0, 0, 0}}},
		{"D", voicing{perString{-1, -1, 0, 2, 3, 1}, perString{0, 0, 1, 3, 4, 2}}, voicing{perString{-1, -1, 0, 2, 3, 1}, perString{0, 0, 0, 2, 3, 1}}},
	},
	"7": {
		{"C", voicing{perString{-1, 3, 2, 3, 1, -1}, perString{0, 3, 2, 4, 1, 0}}, voicing{perString{-1, 3, 2, 3, 1, 0}, perString{0, 3, 2, 4, 1, 0}}},
		{"A", voicing{perString{-1, 0, 2, 0, 2, 0}, perString{0, 1, 3, 1, 4, 1}}, voicing{perString{-1, 0, 2, 0, 2, 0}, perString{0, 0, 2, 0, 3, 0}}},
		{"G", voicing{perString{3, 2, 0, 0, 0, 1}, perString{4, 3, 1, 1, 1, 2}}, voicing{perString{3, 2, 0, 0, 0, 1}, perString{3, 2, 0, 0, 0, 1}}},
		{"E", voicing{perString{0, 2, 0, 1, 0, 0}, perString{1, 3, 1, 2, 1, 1}}, voicing{perString{0, 2, 0, 1, 0, 0}, perString{0, 2, 0, 1, 0, 0}}},
		{"D", voicing{perString{-1, -1, 0, 2, 1, 2}, perString{0, 0, 1, 3, 2, 4}}, voicing{perString{-1, -1, 0, 2, 1, 2}, perString{0, 0, 0, 2, 1, 3}}},
	},
	"m7": {
		{"C", voicing{perString{-1, 3, 1, 3, 1, -1}, perString{0, 3, 1, 4, 1, 0}}, voicing{perString{-1, 3, 1, 3, 1, -1}, perString{0, 3, 1, 4, 1, 0}}},
		{"A", voicing{perString{-1, 0, 2, 0, 1, 0}, perString{0, 1, 3, 1, 2, 1}}, voicing{perString{-1, 0, 2, 0, 1, 0}, perString{0, 0, 2, 0, 1, 0}}},
		{"G", voicing{perString{3, 1, 3, 0, -1, -1}, perString{3, 2, 4, 1, 0, 0}}, voicing{perString{3, 1, 3, 0, -1, -1}, perString{2, 1, 3, 0, 0, 0}}},
		{"E", voicing{perString{0, 2, 0, 0, 0, 0}, perString{1, 3, 1, 1, 1, 1}}, voicing{perString{0, 2, 0, 0, 0, 0}, perString{0, 2, 0, 0, 0, 0}}},
		{"D", voicing{perString{-1, -1, 0, 2, 1, 1}, perString{0, 0, 1, 3, 2, 2}}, voicing{perString{-1, -1, 0, 2, 1, 1}, perString{0, 0, 0, 2, 1, 1}}},
	},
	"maj7": {
		{"C", voicing{perString{-1, 3, 2, 0, 0, 0}, perString{0, 4, 3, 1, 1, 1}}, voicing{perString{-1, 3, 2, 0, 0, 0}, perString{0, 3, 2, 0, 0, 0}}},
		{"A", voicing{perString{-1, 0, 2, 1, 2, 0}, perString{0, 1, 3, 2, 4, 1}}, voicing{perString{-1, 0, 2, 1, 2, 0}, perString{0, 0, 2, 1, 3, 0}}},
		{"G", voicing{perString{3, 2, 0, 0, 0, 2}, perString{4, 2, 1, 1, 1, 3}}, voicing{perString{3, 2, 0, 0, 0, 2}, perString{3, 1, 0, 0, 0, 2}}},
		{"E", voicing{perString{0, 2, 1, 1, 0, 0}, perString{1, 4, 2, 3, 1, 1}}, voicing{perString{0, 2, 1, 1, 0, 0}, perString{0, 3, 1, 2, 0, 0}}},
		{"D", voicing{perString{-1, -1, 0, 2, 2, 2}, perString{0, 0, 1, 3, 3, 3}}, voicing{perString{-1, -1, 0, 2, 2, 2}, perString{0, 0, 0, 1, 2, 3}}},
	},
}

// CagedPositions returns nil when the root is unknown or the suffix has no CAGED shapes.
func CagedPositions(root string, suffix string) []model.ChordsDbGuitarPosition {
	rootPitchClass, isKnownRoot := pitchClasses[root]
	shapes := cagedShapesBySuffix[suffix]
	if !isKnownRoot || len(shapes) == 0 {
		return nil
	}
	positions := make([]model.ChordsDbGuitarPosition, 0, len(shapes))
	for _, shape := range shapes {
		positions = append(positions, shape.positionFor(rootPitchClass))
	}
	return positions
}

func (shape cagedShape) positionFor(rootPitchClass int) model.ChordsDbGuitarPosition {
	shift := semitonesUpFrom(pitchClasses[shape.name], rootPitchClass)
	if shift == 0 {
		return toChordsDbPosition(shape.open.frets, shape.open.fingers)
	}
	return toChordsDbPosition(shiftFrets(shape.movable.frets, shift), shape.movable.fingers)
}

func semitonesUpFrom(fromPitchClass int, toPitchClass int) int {
	return ((toPitchClass-fromPitchClass)%semitones + semitones) % semitones
}

func shiftFrets(frets perString, shift int) perString {
	for i, fret := range frets {
		if fret != mutedString {
			frets[i] = fret + shift
		}
	}
	return frets
}

func toChordsDbPosition(frets perString, fingers perString) model.ChordsDbGuitarPosition {
	baseFret := diagramBaseFret(frets)
	barres := []int{}
	isFullBarre := false
	if barreFret, barredStrings := indexFingerBarre(frets, fingers); len(barredStrings) > 1 {
		barres = append(barres, toDiagramFret(barreFret, baseFret))
		isFullBarre = barredStrings[0] == lowestSoundingString(frets)
	}
	return model.ChordsDbGuitarPosition{
		Frets:    toDiagramFrets(frets, baseFret),
		Fingers:  fingers[:],
		Barres:   barres,
		Capo:     isFullBarre,
		BaseFret: baseFret,
		Midi:     midiNotes(frets),
	}
}

func diagramBaseFret(frets perString) int {
	lowest, highest := 0, 0
	for _, fret := range frets {
		if fret <= openString {
			continue
		}
		if lowest == 0 || fret < lowest {
			lowest = fret
		}
		highest = max(highest, fret)
	}
	if highest <= diagramHeight {
		return 1
	}
	return lowest
}

func toDiagramFrets(frets perString, baseFret int) []int {
	diagramFrets := make([]int, stringCount)
	for i, fret := range frets {
		diagramFrets[i] = toDiagramFret(fret, baseFret)
	}
	return diagramFrets
}

func toDiagramFret(fret int, baseFret int) int {
	if fret <= openString {
		return fret
	}
	return fret - baseFret + 1
}

func fromDiagramFret(diagramFret int, baseFret int) int {
	if diagramFret <= openString {
		return diagramFret
	}
	return diagramFret + baseFret - 1
}

func midiNotes(frets perString) []int {
	notes := []int{}
	for i, fret := range frets {
		if fret != mutedString {
			notes = append(notes, standardTuningMidi[i]+fret)
		}
	}
	return notes
}

func indexFingerBarre(frets perString, fingers perString) (barreFret int, barredStrings []int) {
	for i, fret := range frets {
		if fingers[i] != indexFinger || fret <= openString {
			continue
		}
		if len(barredStrings) == 0 {
			barreFret = fret
		}
		if fret == barreFret {
			barredStrings = append(barredStrings, i)
		}
	}
	return barreFret, barredStrings
}

func lowestSoundingString(frets perString) int {
	for i, fret := range frets {
		if fret != mutedString {
			return i
		}
	}
	return mutedString
}

// AppendNewVoicings keeps existing untouched and appends the generated
// positions it does not already cover, sorted by baseFret.
func AppendNewVoicings(existing []model.ChordsDbGuitarPosition, generated []model.ChordsDbGuitarPosition) []model.ChordsDbGuitarPosition {
	known := make([]perString, 0, len(existing)+len(generated))
	for _, position := range existing {
		known = append(known, fretsOf(position))
	}

	added := []model.ChordsDbGuitarPosition{}
	for _, position := range generated {
		frets := fretsOf(position)
		if anyDiffersOnlyByMutedStrings(known, frets) {
			continue
		}
		known = append(known, frets)
		added = append(added, position)
	}
	sort.SliceStable(added, func(i, j int) bool {
		return added[i].BaseFret < added[j].BaseFret
	})

	return append(append([]model.ChordsDbGuitarPosition{}, existing...), added...)
}

func anyDiffersOnlyByMutedStrings(known []perString, frets perString) bool {
	for _, other := range known {
		if differOnlyByMutedStrings(other, frets) {
			return true
		}
	}
	return false
}

func differOnlyByMutedStrings(a perString, b perString) bool {
	aPlaysOnlyStringsOfB, bPlaysOnlyStringsOfA := true, true
	for i := range a {
		aPlays, bPlays := a[i] != mutedString, b[i] != mutedString
		switch {
		case aPlays && bPlays && a[i] != b[i]:
			return false
		case aPlays && !bPlays:
			aPlaysOnlyStringsOfB = false
		case bPlays && !aPlays:
			bPlaysOnlyStringsOfA = false
		}
	}
	return aPlaysOnlyStringsOfB || bPlaysOnlyStringsOfA
}

func fretsOf(position model.ChordsDbGuitarPosition) perString {
	frets := perString{mutedString, mutedString, mutedString, mutedString, mutedString, mutedString}
	for i := 0; i < len(position.Frets) && i < stringCount; i++ {
		frets[i] = fromDiagramFret(position.Frets[i], position.BaseFret)
	}
	return frets
}
