package generator

import (
	"github.com/demonyze/infernote/internal/model"
)

func Run(runner model.Runner) error {
	chordRocksImport, err := runner.ChordRocksGuitarImporter.Import()
	if err != nil {
		return err
	}

	chordsDbImport, err := runner.ChordsDbGuitarImporter.Import()
	if err != nil {
		return err
	}

	chordTypes := Types(runner.LanguageImport.Types)
	chordsMap := ChordsAsMap(CreateChordParams{
		chordsDbImport,
		chordRocksImport,
		chordTypes,
	})

	infernote := model.Infernote{
		Language: runner.LanguageImport.Language,
		Types:    chordTypes,
		Chords:   chordsMap,
	}

	exportError := runner.InfernoteExporter.Export(infernote)
	if exportError != nil {
		return exportError
	}

	return nil
}
