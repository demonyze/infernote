package generator

import (
	"fmt"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

type CreateChordParams struct {
	chordsDbImport   model.ChordsDbGuitarImport
	chordRocksImport model.ChordRocksGuitarImport
}

func createChord(
	chordsDbChord model.ChordsDbGuitarChord,
	chordsRockChord model.ChordRocksGuitarChord,
	note model.Note,
) model.Chord {
	return model.Chord{
		Id:              utils.CreateId(note.Name, chordsDbChord.Suffix),
		Name:            fmt.Sprintf("%s %s", note.Name, chordsDbChord.Suffix),
		NameShort:       fmt.Sprintf("%s%s", note.Name, utils.AlternateSuffix(chordsDbChord.Suffix)),
		GuitarPositions: chordsDbChord.Positions,
		Root:            note,
		Type:            chordsDbChord.Suffix,
		Notes:           utils.NotesfromStringNames(chordsRockChord.Notes),
	}
}

func ChordsAsMap(params CreateChordParams) map[string]model.Chord {
	chords := make(map[string]model.Chord)
	for _, note := range data.Notes {
		var chordsDbNoteName = mapChordsDbNote(note)
		for _, chordsDbChord := range params.chordsDbImport.Chords[chordsDbNoteName] {
			id := utils.CreateId(note.Name, chordsDbChord.Suffix)
			chordRocksChord := params.chordRocksImport[note.Name][chordsDbChord.Suffix]
			chords[id] = createChord(chordsDbChord, chordRocksChord, note)
		}
	}
	return chords
}

func Types(data map[string]string) map[string]model.ChordType {
	types := make(map[string]model.ChordType)
	for key, value := range data {
		types[key] = model.ChordType{
			Id:   key,
			Name: value,
		}
	}
	return types
}

func mapChordsDbNote(note model.Note) string {
	n := []string{"A#", "D#", "G#"}
	for _, key := range n {
		if note.Name == key {
			return note.AlternativeName
		}
	}
	return note.Name
}
