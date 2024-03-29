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

	chordsMap := ChordsAsMap(CreateChordParams{
		chordsDbImport,
		chordRocksImport,
	})

	types := Types(runner.LanguageImport.Types)

	infernote := model.Infernote{
		Language: runner.LanguageImport.Language,
		Types:    types,
		Chords:   chordsMap,
	}

	exportError := runner.InfernoteExporter.Export(infernote)
	if exportError != nil {
		return exportError
	}

	return nil
}
