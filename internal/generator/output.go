package generator

import (
	"fmt"

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
) model.Chord {
	return model.Chord{
		Id:              utils.CreateId(chordsDbChord.Key, chordsDbChord.Suffix),
		Name:            fmt.Sprintf("%s %s", chordsDbChord.Key, chordsDbChord.Suffix),
		GuitarPositions: chordsDbChord.Positions,
		Root:            model.Note{Name: chordsDbChord.Key},
		Type:            chordsDbChord.Suffix,
		Notes:           utils.NotesfromStringNames(chordsRockChord.Notes),
	}
}

func ChordsAsMap(params CreateChordParams) map[string]model.Chord {
	chords := make(map[string]model.Chord)
	for _, key := range params.chordsDbImport.Keys {
		for _, chordsDbChord := range params.chordsDbImport.Chords[key] {
			id := utils.CreateId(chordsDbChord.Key, chordsDbChord.Suffix)
			chordRocksChord := params.chordRocksImport[key][chordsDbChord.Suffix]
			chords[id] = createChord(chordsDbChord, chordRocksChord)
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
