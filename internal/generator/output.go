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
	languageImports  []model.LanguageImport
}

func createChord(
	chordsDbChord model.ChordsDbGuitarChord,
	chordsRockChord model.ChordRocksGuitarChord,
	note model.Note,
	suffix string,
	languageImports []model.LanguageImport,
) model.Chord {
	lang := make(map[string]model.ChordLang)
	for _, imp := range languageImports {
		typeName := imp.Types[suffix]
		lang[imp.Language] = model.ChordLang{
			Name: fmt.Sprintf("%s %s", note.Name, typeName),
			Type: typeName,
		}
	}
	return model.Chord{
		Id:              utils.CreateId(note.Name, chordsDbChord.Suffix),
		NameShort:       fmt.Sprintf("%s%s", note.Name, utils.AlternateSuffix(chordsDbChord.Suffix)),
		GuitarPositions: chordsDbChord.Positions,
		Root:            note,
		Lang:            lang,
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
			chords[id] = createChord(
				chordsDbChord,
				chordRocksChord,
				note,
				chordsDbChord.Suffix,
				params.languageImports,
			)
		}
	}
	return chords
}

func Types(imports []model.LanguageImport) map[string]model.ChordType {
	types := make(map[string]model.ChordType)
	for _, imp := range imports {
		for id, name := range imp.Types {
			ct, exists := types[id]
			if !exists {
				ct = model.ChordType{Id: id, Lang: make(map[string]string)}
			}
			ct.Lang[imp.Language] = name
			types[id] = ct
		}
	}
	return types
}

func mapChordsDbNote(note model.Note) string {
	switch note.Name {
	case "A#":
		return note.AlternativeName // "Bb"
	case "D#":
		return note.AlternativeName // "Eb"
	case "G#":
		return note.AlternativeName // "Ab"
	case "C#":
		return "Csharp"
	case "F#":
		return "Fsharp"
	}
	return note.Name
}
