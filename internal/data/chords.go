package data

import (
	"fmt"

	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

type Structure = string

const (
	AsArray Structure = "asArray"
	AsMap   Structure = "asMap"
)

func createChord(chordImport model.ChordsDbGuitarChord) model.Chord {
	return model.Chord{
		Id:              utils.CreateId(chordImport.Key, chordImport.Suffix),
		Name:            fmt.Sprintf("%s %s", chordImport.Key, chordImport.Suffix),
		GuitarPositions: chordImport.Positions,
		Root:            model.Note{Name: chordImport.Key},
		Type:            chordImport.Suffix,
	}
}

func ChordsAsMap(dataImport model.ChordsDbGuitarImport) map[string]model.Chord {
	chords := make(map[string]model.Chord)

	for _, key := range dataImport.Keys {
		for _, chordImport := range dataImport.Chords[key] {
			id := utils.CreateId(chordImport.Key, chordImport.Suffix)
			chords[id] = createChord(chordImport)
		}
	}

	return chords
}

func ChordsAsArray(dataImport model.ChordsDbGuitarImport) []model.Chord {
	var chords []model.Chord

	for _, key := range dataImport.Keys {
		for _, chordImport := range dataImport.Chords[key] {
			chords = append(chords, createChord(chordImport))
		}
	}

	return chords
}
